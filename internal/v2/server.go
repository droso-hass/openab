package v2

import (
	"log/slog"
	"net/http"

	"github.com/droso-hass/openab/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// http://192.168.1.132/vl
// locate.jsp ?
// /vl/bc.jsp
// /vl/rfid.jsp

// platform: bootcode, locate
// broad: sounds, choreos
// ping: rfid, recordings

func Setup(r *chi.Mux) {
	r.Mount("/vl/bc.jsp", test())
	r.Mount("/vl/locate.jsp", locate())
	r.Mount("/vl/hello", hello())
}

func test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		slog.Info("new connection from v2", "version", q.Get("v"), "mac", q.Get("m"))
		utils.SendFile(w, r, "./firmware/v2/nominal.bin", "application/octet-stream")
	}
}

func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, 200)
		render.PlainText(w, r, "http hello world !")
	}
}

func locate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, 200)
		render.PlainText(w, r, "ping 192.168.1.21\nbroad 192.168.1.21\nxmpp_domain dev.nab.lan:5000")
	}
}
