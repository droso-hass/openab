package api

import (
	"encoding/json"
	"strings"

	"github.com/droso-hass/openab/internal/common"
	"github.com/nats-io/nats.go"
)

func (a *API) processSub(m *nats.Msg) {
	for _, r := range a.receivers {
		sp := strings.Split(m.Subject, ".")
		if sp[2] == "led" {
			c := common.NabLedCmd{}
			err := json.Unmarshal(m.Data, &c)
			if err == nil {
				r.Led(sp[1], c)
			}
		}
	}
}
