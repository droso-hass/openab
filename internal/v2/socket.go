package v2

import (
	"net"
)

type NabConn struct {
	addr string
	conn net.Conn
	stop bool
}

func New(addr string) *NabConn {
	return &NabConn{
		addr: addr,
		stop: false,
	}
}

func (n *NabConn) Disconnect() error {
	n.stop = true
	return n.conn.Close()
}

func (n *NabConn) Connect() error {
	conn, err := net.Dial("tcp", n.addr)
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
			_, err := conn.Read(buf)
			if err == nil {
				n.processNabMessage(buf)
			}
		}
	}()

	return nil
}
