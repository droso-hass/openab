package v2

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/droso-hass/openab/internal/udp"
	"github.com/droso-hass/openab/internal/utils"
	"github.com/go-chi/chi/v5"
)

// http://192.168.1.132/vl
// locate.jsp ?
// /vl/bc.jsp
// /vl/rfid.jsp

// platform: bootcode, locate
// broad: sounds, choreos
// ping: rfid, recordings

var conns = map[string]*NabConn{}
var v2chan chan udp.UDPPacket

func Init(r *chi.Mux) {
	r.Mount("/vl/bc.jsp", bootcode())
	v2chan = make(chan udp.UDPPacket)
	go handleUDP(v2chan)
	/*
	   c := *New("127.0.0.1", "")
	   e := c.Connect()

	   	if e != nil {
	   		log.Fatal(e)
	   	}

	   conns["127.0.0.1"] = &c
	   debug("127.0.0.1")
	*/
}

func fftest() {
	ch := make(chan []byte)
	err := convertPlayer("./static/lisa.mp3", ch)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("./static/lisanab.mp3", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	for {
		x := <-ch
		if x == nil {
			break
		} else if len(x) > 0 {
			_, err = f.Write(x)
			if err != nil {
				log.Fatal(err)
			}
			//time.Sleep(time.Microsecond * 100)
		}
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func debug(m string) {
	/*e := conns[m].write("07;2;255")
	if e != nil {
		log.Fatal(e)
	}*/

	/*e := conns[m].write("07;1")
	if e != nil {
		log.Fatal(e)
	}*/
	e := conns[m].write("00;ping")
	if e != nil {
		log.Fatal(e)
	}

	uaddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:4000", conns[m].ip))
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan []byte)
	err = convertPlayer("./server/static/lisa.mp3", ch)
	if err != nil {
		log.Fatal(err)
	}
	for {
		x := <-ch
		if x == nil {
			break
		} else if l := len(x); l > 0 {
			if conns[m].playWritten+l > 1024 {
				fmt.Println("stop")
				conns[m].playMtx.Lock()
				conns[m].playWritten = 0
			}
			udp.Write(udp.UDPPacket{
				Addr: uaddr,
				Type: udp.UDPTypeSound,
				Data: x,
			})
			conns[m].playWritten += l
			fmt.Println("udpw")
			//time.Sleep(time.Second * 1)
		}
	}
	fmt.Println("ok")
}

func bootcode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./server/static/nominal.bin", "application/octet-stream")
		//utils.SendFile(w, r, "./static/bc.jsp", "application/octet-stream")

		go func() {
			time.Sleep(time.Second * 3)

			ip := utils.GetIPFromRequest(r)
			slog.Info("connecting to " + ip)

			// if a connection is already open, close it
			c, ok := conns[ip]
			if ok {
				c.Disconnect()
				c.Connect()
			} else {
				c := *New(ip, q.Get("m"))
				err := c.Connect()
				if err != nil {
					log.Fatal(err)
				}
				conns[ip] = &c
			}

			udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:4000", ip))
			if err != nil {
				slog.Error("error parsing udp address", utils.ErrAttr(err))
			} else {
				udp.RegisterCallback(udpAddr, v2chan)
			}

			time.Sleep(time.Second * 2)
			//initseq(ip)
			debug(ip)

			reader := bufio.NewReader(os.Stdin)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				err = conns[ip].write(str)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
}

func initseq(m string) {
	green_breath := "03;4;0;00FF00;100;00EE00;100;00DD00;100;00CC00;100;00BB00;100;00AA00;100;009900;100;008800;100;007700;100;006600;100;005500;100;004400;100;003300;100;002200;100;001100;100;000000;100;001100;100;002200;100;003300;100;004400;100;005500;100;006600;100;007700;100;008800;100;009900;100;00AA00;100;00BB00;100;00CC00;100;00DD00;100;00EE00;100"
	conns[m].write(green_breath + "\n02;0;0;0;1;0\n02;1;0;0;1;0")
	/*time.Sleep(time.Second * 2)
	conns[m].write("06;1")
	time.Sleep(time.Second * 7)
	conns[m].write("06;0")
	conns[m].recFile.Close()
	fmt.Println("ok")*/
}
