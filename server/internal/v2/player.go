package v2

import (
	"errors"
	"fmt"
	"time"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/udp"
)

var ErrBufferFull = errors.New("player buffer is full")

func (n *NabConn) playStream(data []byte) error {
	x, err := convertPlayerChunk(data)
	if err != nil {
		return err
	}
	select {
	case n.playChan <- x:
		return nil
	default:
		return ErrBufferFull
	}
}

func (n *NabConn) playLink(url string) error {
	// make sure channel is empty
	for i := len(n.playChan); i > 0; i-- {
		<-n.playChan
	}
	return convertPlayerFile(url, n.playChan, 1024)
}

func (n *NabConn) playLoop() {
	var lastPacket []byte = nil
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
				x := <-n.playChan
				if x == nil {
					// everything has been sent
					time.Sleep(100 * time.Millisecond)
					if !n.isPlaying.Load() {
						// if not yet playing, force start the player
						n.write("07;1")
						n.isPlaying.Store(true)
						n.pub.PlayerState(n.mac, common.NabAudioEvent{State: common.NabAudioRunning})
					}
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
}
