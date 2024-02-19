package v2

import (
	"encoding/hex"
	"fmt"
	"log/slog"
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
		fmt.Printf("%+v\n", data)
		if data.Type == udp.UDPTypeSound {
			err := nab.handleRecording(string(data.Data))
			if err != nil {
				slog.Warn("V2: error processing recording", utils.ErrAttr(err))
			}
		}
	}
}

func (n *NabConn) handleRecording(rawdata string) error {
	if len(n.recData) >= RecDataSize {
		wav, err := makeWav(n.recData)
		if err != nil {
			return err
		}
		filedata, err := convertRecording(wav)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("./rec/%s_%d.wav", strings.Replace(n.mac, ":", "", -1), time.Now().Unix())
		err = utils.WriteFile(filename, filedata)
		if err != nil {
			return err
		}
		n.recData = []byte{}
	}
	h, err := hex.DecodeString(rawdata)
	if err != nil {
		return err
	}
	n.recData = append(n.recData, h...)
	return nil
}
