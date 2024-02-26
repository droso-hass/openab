package v2

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
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
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./server/static/nominal.bin", "application/octet-stream")

		go func() {
			time.Sleep(time.Second * 3)

			ip := utils.GetIPFromRequest(r)
			slog.Info("connecting to " + ip)

			// if a connection is already open, close it
			c, ok := n.conns[ip]
			if ok {
				c.Disconnect()
				c.Connect()
			} else {
				c := *NewNab(ip, q.Get("m"), n.pub)
				err := c.Connect()
				if err != nil {
					log.Fatal(err)
				}
				n.conns[ip] = &c
			}

			udp.RegisterCallback(c.udpAddr, n.udpChan)

			time.Sleep(time.Second * 2)
			//n.initseq(ip)

			reader := bufio.NewReader(os.Stdin)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				err = n.conns[ip].write(str)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
}

func (n *NabV2) initseq(m string) {
	green_breath := "03;4;0;00FF00;100;00EE00;100;00DD00;100;00CC00;100;00BB00;100;00AA00;100;009900;100;008800;100;007700;100;006600;100;005500;100;004400;100;003300;100;002200;100;001100;100;000000;100;001100;100;002200;100;003300;100;004400;100;005500;100;006600;100;007700;100;008800;100;009900;100;00AA00;100;00BB00;100;00CC00;100;00DD00;100;00EE00;100"
	n.conns[m].write(green_breath + "\n02;0;0;0;1;0\n02;1;0;0;1;0")
	/*time.Sleep(time.Second * 2)
	conns[m].write("06;1")
	time.Sleep(time.Second * 7)
	conns[m].write("06;0")
	conns[m].recFile.Close()
	fmt.Println("ok")*/
}
