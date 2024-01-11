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

func SetupRoutes(r *chi.Mux) {
	r.Mount("/vl/bc.jsp", bootcode())
	//connect("127.0.0.1:5000")
	//connect("10.20.5.2:5000")
}

func bootcode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./firmware/v2/nominal.bin", "application/octet-stream")
		go func() {
			time.Sleep(time.Second * 3)
			connect(utils.GetIPFromRequest(r) + ":5000")
		}()
	}
}
