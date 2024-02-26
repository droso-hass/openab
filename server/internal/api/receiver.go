package api

import (
	"encoding/json"
	"strings"

	"github.com/droso-hass/openab/internal/common"
	"github.com/nats-io/nats.go"
)

func (a *API) getReceiver(mac string) common.INabReceiver {
	for _, h := range a.receivers {
		recv := h.FindReceiver(mac)
		if recv != nil {
			return recv
		}
	}
	return nil
}

func (a *API) processSub(m *nats.Msg) {
	sp := strings.Split(m.Subject, ".")
	recv := a.getReceiver(sp[1])
	if recv == nil {
		return
	}

	if sp[2] == "led" {
		c := common.NabLedCmd{}
		err := json.Unmarshal(m.Data, &c)
		if err == nil {
			recv.Led(sp[1], c)
		}
	}
}
