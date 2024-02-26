package v2

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/udp"
	"github.com/droso-hass/openab/internal/utils"
	"github.com/go-chi/chi/v5"
)

type NabV2 struct {
	pub     common.INabSender
	conns   map[string]*NabConn
	udpChan chan udp.UDPPacket
}

func New(r *chi.Mux, snd common.INabSender) *NabV2 {
	v2chan := make(chan udp.UDPPacket)
	n := NabV2{pub: snd, conns: map[string]*NabConn{}, udpChan: v2chan}
	go n.handleUDP(v2chan)
	r.Mount("/vl/bc.jsp", bootcode(&n))
	return &n
}

func (n *NabV2) FindReceiver(mac string) common.INabReceiver {
	for _, v := range n.conns {
		if v.mac == mac {
			return v
		}
	}
	return nil
}

func bootcode(n *NabV2) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		mac := q.Get("m")
		ip := utils.GetIPFromRequest(r)
		version := q.Get("v")
		slog.Info("new connection from v2", "version", version, "mac", mac)
		utils.SendFile(w, r, "./server/static/nominal.bin", "application/octet-stream")
		go n.connectNab(mac, ip, version)
	}
}

func (n *NabV2) connectNab(mac string, ip string, fwver string) {
	time.Sleep(time.Millisecond * 1500)

	slog.Info("connecting to " + ip)

	// if a connection is already open, close it
	c, ok := n.conns[ip]
	if ok {
		c.Disconnect()
		c.Connect()
	} else {
		c := *NewNab(ip, mac, n.pub)
		err := c.Connect()
		if err != nil {
			log.Fatal(err)
		}
		n.conns[ip] = &c
	}
	udp.RegisterCallback(c.udpAddr, n.udpChan)
	go c.playLoop()

	n.pub.Status(mac, common.NabStatus{
		Connected: true,
		IP:        ip,
		HWVersion: common.NabVersionTagTag,
		FWVersion: fwver,
		SWVersion: "TODO: return server version",
	})

	// TODO: add ping/pong to detect disconnects
}
