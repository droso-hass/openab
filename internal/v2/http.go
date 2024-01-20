package v2

import (
	"log/slog"
	"net/http"
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
	c.Connect()
	c.write("031010100FF001000")*/
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

			time.Sleep(time.Second * 2)
			conns[q.Get("m")].write("030100000FF003000")
			time.Sleep(time.Second * 1)
			conns[q.Get("m")].write("030010000FF003000")
			time.Sleep(time.Second * 1)
			conns[q.Get("m")].write("030001000FF003000")
		}()
	}
}
