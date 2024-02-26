package api

import (
	"log/slog"
	"os"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/config"
	"github.com/nats-io/nats.go"
)

type API struct {
	nc        *nats.Conn
	ec        *nats.EncodedConn
	receivers []common.INabReceiverHander
	syncCmds  map[string]common.NabSyncedItems
}

func New(url string) *API {
	opts := []nats.Option{
		nats.Name(config.ConfigData.ClientID),
	}
	if config.ConfigData.NatsUsername != "" && config.ConfigData.NatsPassword != "" {
		opts = append(opts, nats.UserInfo(config.ConfigData.NatsUsername, config.ConfigData.NatsPassword))
	}
	nc, err := nats.Connect(config.ConfigData.NatsServer, opts...)

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
		syncCmds:  map[string]common.NabSyncedItems{},
	}
}

func (a *API) Listen(receivers []common.INabReceiverHander) {
	a.receivers = receivers
	a.nc.Subscribe("openab.>", a.processSub)
}

func (a *API) Stop() {
	a.nc.Drain()
}
