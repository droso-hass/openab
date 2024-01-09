package v2

import "net"

type nab struct {
	conn net.Conn
	mac  string
}
