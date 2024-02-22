package v2

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/droso-hass/openab/internal/udp"
	"github.com/droso-hass/openab/internal/utils"
)

func handleUDP(ch chan udp.UDPPacket) {
	for {
		data := <-ch
		nab, ok := conns[data.Addr.IP.String()]
		if !ok {
			slog.Warn("V2: udp packet received but no handler is associated with this ip", "addr", data.Addr)
			continue
		}
		if data.Type == udp.UDPTypeSoundData {
			err := nab.handleRecording(data.Data)
			if err != nil {
				slog.Warn("V2: error processing recording", utils.ErrAttr(err))
			}
		} else if data.Type == udp.UDPTypeSoundSend {
			// allow next packet
			i, err := strconv.ParseUint(string(data.Data), 10, 8)
			if err == nil {
				ii := uint8(i)
				if ii != nab.playLastAck {
					nab.playLastAck = ii
					fmt.Printf("ack: %d\n", nab.playLastAck)
					if nab.playMtxLocked {
						nab.playMtxLocked = false
						nab.playMtx.Unlock()
					}
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}

func (n *NabConn) handleRecording(rawdata []byte) error {
	if len(n.recData) >= RecDataSize {
		wav, err := makeWav(n.recData)
		if err != nil {
			return err
		}
		filedata, err := convertRecording(wav)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("./server/rec/%s_%d.wav", strings.Replace(n.mac, ":", "", -1), time.Now().Unix())
		err = utils.WriteFile(filename, filedata)
		if err != nil {
			return err
		}
		n.recData = []byte{}
	}
	n.recData = append(n.recData, rawdata...)
	return nil
}
