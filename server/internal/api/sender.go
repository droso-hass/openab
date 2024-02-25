package api

import (
	"encoding/json"
	"log/slog"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/utils"
)

func (a *API) pub(mac string, topic string, data []byte) {
	err := a.nc.Publish("openab."+mac+"."+topic, data)
	if err != nil {
		slog.Warn("error publishing data", utils.ErrAttr(err))
	}
}

func (a *API) Ear(mac string, data common.NabEarEvent) {

}

func (a *API) Status(mac string, data common.NabStatus) {
	d, err := json.Marshal(data)
	if err != nil {
		slog.Error("error marshaling json", utils.ErrAttr(err))
	} else {
		a.pub(mac, "status", d)
	}
}

func (a *API) Rfid(mac string, data string) {

}

func (a *API) PlayerState(mac string, data common.NabAudio) {

}

func (a *API) Button(mac string, data common.NabButton, value uint8) {

}

func (a *API) RecorderData(mac string, data []byte) {

}
