package api

import (
	"log/slog"
	"os"

	"github.com/droso-hass/openab/internal/common"
	"github.com/nats-io/nats.go"
)

type API struct {
	nc        *nats.Conn
	ec        *nats.EncodedConn
	receivers []common.INabReceiverHander
}

func New(url string) *API {
	nc, err := nats.Connect(url, nats.Name("openab server"))
	if err != nil {
		slog.Error("failed to connect to the NATS server")
		os.Exit(1)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		slog.Error("failed to create NATS Encoded connection")
		os.Exit(1)
	}
	return &API{
		nc:        nc,
		ec:        ec,
		receivers: []common.INabReceiverHander{},
	}
}

func (a *API) Listen(receivers []common.INabReceiverHander) {
	a.receivers = receivers
	a.nc.Subscribe("openab.>", a.processSub)
}

func (a *API) Stop() {
	a.nc.Drain()
}
