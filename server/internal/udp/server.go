package udp

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"sync"

	"github.com/droso-hass/openab/internal/utils"
)

var ErrCannotWrite = errors.New("unable to write data")

var server UDPServer

func Start(port int) error {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("unable to resolve udp address: %s", err.Error())
		return err
	}
	conn4, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatalf("unable to initialize UDPConn: %s", err.Error())
		return err
	}

	server = UDPServer{
		conn4:         conn4,
		callbacks:     map[string]chan UDPPacket{},
		callbacksLock: sync.Mutex{},
	}

	go listen()

	return nil
}

func listen() {
	for {
		buffer := make([]byte, 520)
		recvSize, addr, err := server.conn4.ReadFrom(buffer)
		if err != nil {
			slog.Debug("error receiving udp packet", utils.ErrAttr(err))
		}
		processData(recvSize, addr, buffer)
	}
}

func processData(recvSize int, addr net.Addr, buffer []byte) {
	udpAddr, err := net.ResolveUDPAddr(addr.Network(), addr.String())
	if err != nil {
		slog.Error("error converting to udp address", utils.ErrAttr(err))
		return
	}
	if recvSize == 520 {
		slog.Warn("received udp message is the same size as the buffer, it may have been truncated")
	}

	server.callbacksLock.Lock()
	defer server.callbacksLock.Unlock()
	channel, ok := server.callbacks[udpAddr.String()]
	if ok {
		channel <- UDPPacket{
			Addr: udpAddr,
			Type: string(buffer[0:3]),
			Data: buffer[3:recvSize],
		}
	} else {
		slog.Debug("no callback registered for this address", "addr", udpAddr)
	}
}

func Write(pkt UDPPacket) error {
	_, err := server.conn4.WriteTo(append([]byte(pkt.Type), pkt.Data...), pkt.Addr)
	if err != nil {
		slog.Warn("error writing to udp4", utils.ErrAttr(err))
		return ErrCannotWrite
	}
	return nil
}

func RegisterCallback(addr *net.UDPAddr, ch chan UDPPacket) {
	server.callbacksLock.Lock()
	defer server.callbacksLock.Unlock()
	server.callbacks[addr.String()] = ch
}
