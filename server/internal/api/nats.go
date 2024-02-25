package api

import (
	"log/slog"
	"os"

	"github.com/droso-hass/openab/internal/common"
	"github.com/nats-io/nats.go"
)

type API struct {
	nc        *nats.Conn
	receivers []common.INabReceiver
}

func New(url string) *API {
	nc, err := nats.Connect(url, nats.Name("openab server"))
	if err != nil {
		slog.Error("failed to connect to the NATS server")
		os.Exit(1)
	}
	return &API{
		nc:        nc,
		receivers: []common.INabReceiver{},
	}
}

func (a *API) Listen(receivers []common.INabReceiver) {
	a.receivers = receivers
	a.nc.Subscribe("openab.>", a.processSub)
}

func (a *API) Stop() {
	a.nc.Drain()
}
