package v2

import (
	"fmt"
	"log"
	"time"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/udp"
)

func (n *NabConn) write(data string) error {
	_, err := n.conn.Write([]byte(data))
	return err
}

func (n *NabConn) Play(url string) {
	ch := make(chan []byte)
	err := convertPlayer(url, ch, 1024)
	if err != nil {
		log.Fatal(err)
	}
	var lastPacket []byte = nil
out:
	for {
		n.playMtx.Lock()
		n.playMtxLocked = true

		if lastPacket != nil && n.playLastSent > n.playLastAck {
			// if lost, replay last packet
			p := []byte(fmt.Sprintf("%03d", n.playLastSent))
			udp.Write(udp.UDPPacket{
				Addr: n.udpAddr,
				Type: udp.UDPTypeSoundData,
				Data: append(p, lastPacket...),
			})
		} else {
			n.playLastSent++
			p := []byte(fmt.Sprintf("%03d", n.playLastSent))
			for {
				x := <-ch
				if x == nil {
					break out
				} else if l := len(x); l > 0 {
					lastPacket = x
					udp.Write(udp.UDPPacket{
						Addr: n.udpAddr,
						Type: udp.UDPTypeSoundData,
						Data: append(p, x...),
					})
					break
				}
			}
		}
	}

	// everything has been sent
	time.Sleep(100 * time.Millisecond)
	if !n.isPlaying.Load() {
		// if not yet playing, force start the player
		n.write("07;1")
		n.isPlaying.Store(true)
		n.pub.PlayerState(n.mac, common.NabAudioEvent{State: common.NabAudioRunning})
	}
}
