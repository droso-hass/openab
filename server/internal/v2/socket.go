package v2

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"

	"github.com/droso-hass/openab/internal/common"
)

type NabConn struct {
	ip      string
	mac     string
	udpAddr *net.UDPAddr
	conn    net.Conn
	stop    bool
	recData []byte
	pub     common.INabSender
	// player
	playMtx       sync.Mutex
	playMtxLocked bool
	playLastSent  uint8
	playLastAck   uint8
	isPlaying     atomic.Bool
}

func NewNab(ip string, mac string, pub common.INabSender) *NabConn {
	uaddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:4000", ip))
	if err != nil {
		log.Fatal(err)
	}
	n := NabConn{
		ip:            ip,
		mac:           mac,
		udpAddr:       uaddr,
		stop:          false,
		recData:       []byte{},
		pub:           pub,
		playMtx:       sync.Mutex{},
		playMtxLocked: false,
		playLastSent:  0,
		playLastAck:   0,
		isPlaying:     atomic.Bool{},
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
