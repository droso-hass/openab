package v2

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/udp"
	"github.com/droso-hass/openab/internal/utils"
)

func (n *NabConn) stopPlayer() {
	n.isPlaying.Store(false)
	if n.playChanIn != nil {
		n.playChanIn <- nil
	}
}

func (n *NabConn) playStream(data []byte) error {
	if n.playChanIn == nil {
		if n.isPlaying.Load() {
			return common.ErrAlreadyPlaying
		}
		n.playChanOut = make(chan []byte)
		n.playChanIn = make(chan []byte, 100)
		go n.playLoop()
		return utils.FFconvertChunk(utils.FFpcmToMp3, &n.playChanIn, n.playChanOut, 1024, 1*time.Second)
	} else {
		select {
		case n.playChanIn <- data:
			return nil
		default:
			return common.ErrBufferFull
		}
	}
}

func (n *NabConn) playLink(url string) error {
	if n.playChanIn == nil {
		if n.isPlaying.Load() {
			return common.ErrAlreadyPlaying
		}
		n.playChanIn = make(chan []byte)
		n.playChanOut = make(chan []byte)
		cmd := []string{}
		cmd = append(cmd, utils.FFfile...)
		cmd = append(cmd, url)
		cmd = append(cmd, utils.FFtoMP3...)
		go n.playLoop()
		return utils.FFconvertFile(cmd, &n.playChanIn, n.playChanOut, 1024)
	}
	return common.ErrAlreadyPlaying
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
			slog.Debug("loop")
			for {
				x := <-n.playChanOut
				if x == nil {
					// everything has been sent
					slog.Debug("all player data sent")
					if !n.isPlaying.Load() {
						// if not yet playing, force start the player
						slog.Debug("player is not running, starting it")
						n.write("07;1")
						n.isPlaying.Store(true)
						n.pub.PlayerState(n.mac, common.NabAudioEvent{State: common.NabAudioRunning})
					}
					return
				} else if l := len(x); l > 0 {
					slog.Debug("sending audio packet")
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
