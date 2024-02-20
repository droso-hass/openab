package v2

import (
	"net"
)

type NabConn struct {
	ip      string
	mac     string
	conn    net.Conn
	stop    bool
	recData []byte
}

func New(ip string, mac string) *NabConn {
	return &NabConn{
		ip:      ip,
		mac:     mac,
		stop:    false,
		recData: []byte{},
	}
}

func (n *NabConn) Disconnect() error {
	n.stop = true
	return n.conn.Close()
}

func (n *NabConn) Connect() error {
	conn, err := net.Dial("tcp", n.ip+":5000")
	if err != nil {
		return err
	}
	n.stop = false
	n.conn = conn

	// reading loop
	go func() {
		for !n.stop {
			//fmt.Println("reading ...")
			buf := make([]byte, 2048)
			nb, err := conn.Read(buf)
			if err == nil {
				n.processNabMessage(buf[0:nb])
			}
		}
	}()

	return nil
}
