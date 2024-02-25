package v2

import (
	"fmt"
	"strconv"

	"github.com/droso-hass/openab/internal/common"
)

var RecDataSize = 4096

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	fmt.Println(sdata)
	switch sdata[0:2] {
	case "04":
		if sdata[3:4] == "1" {
			n.pub.Button(n.mac, common.NabButtonShort, 0)
		} else if sdata[3:4] == "2" {
			n.pub.Button(n.mac, common.NabButtonDouble, 0)
		} else if sdata[3:4] == "3" {
			n.pub.Button(n.mac, common.NabButtonLongStart, 0)
		} else if sdata[3:4] == "4" {
			n.pub.Button(n.mac, common.NabButtonLongEnd, 0)
		}
	case "05":
		n.pub.Rfid(n.mac, sdata[3:])
	case "07":
		if sdata[3:4] == "1" {
			n.isPlaying.Store(true)
			n.pub.PlayerState(n.mac, common.NabAudioRunning)
		} else if sdata[3:4] == "0" {
			n.pub.PlayerState(n.mac, common.NabAudioStopped)
		}
	case "09":
		i, err := strconv.ParseUint(sdata[3:], 10, 8)
		if err == nil {
			n.pub.Button(n.mac, common.NabButtonVolume, uint8(i))
		}
	case "10":
		tp, err := strconv.ParseUint(sdata[3:4], 10, 8)
		if err == nil {
			pos, err := strconv.ParseUint(sdata[5:], 10, 8)
			if err == nil {
				n.pub.Ear(n.mac, common.NabEarEvent{
					Position: uint8(pos),
					ID:       common.NabEar(tp),
				})
			}
		}
	}
}
