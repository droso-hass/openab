package v2

import (
	"fmt"
	"strconv"

	"github.com/droso-hass/openab/internal/common"
)

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	fmt.Println(sdata)
	switch sdata[0:2] {
	case "04":
		if sdata[3:4] == "1" {
			n.pub.Button(n.mac, common.NabButtonEvent{ID: common.NabButtonShort})
		} else if sdata[3:4] == "2" {
			n.pub.Button(n.mac, common.NabButtonEvent{ID: common.NabButtonDouble})
		} else if sdata[3:4] == "3" {
			n.pub.Button(n.mac, common.NabButtonEvent{ID: common.NabButtonLongStart})
		} else if sdata[3:4] == "4" {
			n.pub.Button(n.mac, common.NabButtonEvent{ID: common.NabButtonLongEnd})
		}
	case "05":
		n.pub.Rfid(n.mac, common.NabRFIDEvent{Value: sdata[3:]})
	case "07":
		if sdata[3:4] == "1" {
			n.isPlaying.Store(true)
			n.pub.PlayerState(n.mac, common.NabAudioEvent{State: common.NabAudioRunning})
		} else if sdata[3:4] == "0" {
			n.isPlaying.Store(false)
			n.pub.PlayerState(n.mac, common.NabAudioEvent{State: common.NabAudioStopped})
		}
	case "09":
		i, err := strconv.ParseUint(sdata[3:], 10, 8)
		if err == nil {
			n.pub.Button(n.mac, common.NabButtonEvent{ID: common.NabButtonVolume, Value: uint8(i)})
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
