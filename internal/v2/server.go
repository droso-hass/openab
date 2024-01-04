package v2

import (
	"log/slog"
	"net/http"

	"github.com/droso-hass/openab/internal/utils"
	"github.com/go-chi/chi/v5"
)

// locate.jsp ?
// /vl/bc.jsp
// /vl/rfid.jsp

// platform: bootcode, locate
// broad: sounds, choreos
// ping: rfid, recordings

func Setup(r *chi.Mux) {
	r.Mount("/vl/bc.jsp", test())
}

func test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./firmware/v2/nominal.bin", "application/octet-stream")
	}
}
