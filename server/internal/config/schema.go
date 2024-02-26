package config

import (
	"flag"
	"os"

	"github.com/droso-hass/openab/internal/utils"
	"github.com/jamiealquiza/envy"
)

var ConfigData Config

type Config struct {
	NatsServer   string
	NatsUsername string
	NatsPassword string
	ClientID     string
}

func Parse(path string) {
	hn, _ := os.Hostname()

	level := flag.String("log-level", "info", "log level")
	server := flag.String("nats-server", "nats://localhost:4222", "nats address, in the format: [tls,nats]://host:port")
	user := flag.String("nats-user", "", "username for the nats server")
	pass := flag.String("nats-password", "", "password for the nats server")
	clientid := flag.String("client-id", hn, "identifier for this device")

	envy.Parse("OPN")
	flag.Parse()

	utils.SetupLogs(*level)

	ConfigData = Config{
		NatsServer:   *server,
		NatsUsername: *user,
		NatsPassword: *pass,
		ClientID:     *clientid,
	}
}
