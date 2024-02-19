package udp

import (
	"net"
	"sync"
)

type UDPServer struct {
	conn4         *net.UDPConn
	callbacks     map[string]chan UDPPacket
	callbacksLock sync.Mutex
}

type UDPPacket struct {
	Data []byte
	Type string
	Addr *net.UDPAddr
}

const (
	UDPTypeSound = "snd"
)
