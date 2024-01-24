package v2

import (
	"bufio"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

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

func SetupRoutes(r *chi.Mux) {
	r.Mount("/vl/bc.jsp", bootcode())
	/*c := New("127.0.0.1:5000")
	e := c.Connect()
	if e != nil {
		log.Fatal(e)
	}
	//c.write("03;4;0;00ff00;100;00f200;100")
	c.write("02;1;0;7;1;0")
	// 02;0;0;8;1;0 02;1;0;7;1;0
	*/
}

func bootcode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./firmware/v2/nominal.bin", "application/octet-stream")

		go func() {
			addr := utils.GetIPFromRequest(r) + ":5000"
			time.Sleep(time.Second * 3)
			slog.Info("connecting to " + addr)

			// if a connection is already open, close it
			c, ok := conns[q.Get("m")]
			if ok {
				c.Disconnect()
				c.Connect()
			} else {
				c := *New(addr)
				c.Connect()
				conns[q.Get("m")] = &c
			}

			reader := bufio.NewReader(os.Stdin)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				conns[q.Get("m")].write(str)
			}
			/*
				time.Sleep(time.Second * 2)
				green_breath := "03;4;0;00FF00;100;00EE00;100;00DD00;100;00CC00;100;00BB00;100;00AA00;100;009900;100;008800;100;007700;100;006600;100;005500;100;004400;100;003300;100;002200;100;001100;100;000000;100;001100;100;002200;100;003300;100;004400;100;005500;100;006600;100;007700;100;008800;100;009900;100;00AA00;100;00BB00;100;00CC00;100;00DD00;100;00EE00;100"
				conns[q.Get("m")].write(green_breath)
				time.Sleep(time.Second * 1)
				conns[q.Get("m")].write("030010000FF003000")
				time.Sleep(time.Second * 1)
				conns[q.Get("m")].write("030001000FF003000")*/
		}()
	}
}
