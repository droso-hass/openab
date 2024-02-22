package v2

import (
	"net"
	"sync"
)

type NabConn struct {
	ip            string
	mac           string
	conn          net.Conn
	stop          bool
	recData       []byte
	playMtx       sync.Mutex
	playMtxLocked bool
	playLastSent  uint8
	playLastAck   uint8
}

func New(ip string, mac string) *NabConn {
	n := NabConn{
		ip:            ip,
		mac:           mac,
		stop:          false,
		recData:       []byte{},
		playMtx:       sync.Mutex{},
		playMtxLocked: false,
		playLastSent:  0,
		playLastAck:   0,
	}
	return &n
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
