package api

import (
	"fmt"

	"github.com/droso-hass/openab/internal/common"
)

func getTopic(mac string, t string) string {
	return fmt.Sprintf("openab.%s.%s", mac, t)
}

func (a *API) Ear(mac string, data common.NabEarEvent) {
	a.ec.Publish(getTopic(mac, "ear.user"), data)
}

func (a *API) Status(mac string, data common.NabStatus) {
	a.ec.Publish(getTopic(mac, "status"), data)
}

func (a *API) Rfid(mac string, data common.NabRFIDEvent) {
	a.ec.Publish(getTopic(mac, "rfid"), data)
}

func (a *API) PlayerState(mac string, data common.NabAudioEvent) {
	a.ec.Publish(getTopic(mac, "player.state"), data)
}

func (a *API) Button(mac string, data common.NabButtonEvent) {
	a.ec.Publish(getTopic(mac, "button"), data)
}

func (a *API) RecorderData(mac string, data []byte) {
	a.nc.Publish(getTopic(mac, "recorder.data"), data)
}
