package utils

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func ServeStatic(r chi.Router, serverRoute string, pathToStaticFolder http.FileSystem) {
	if strings.ContainsAny(serverRoute, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if serverRoute != "/" && serverRoute[len(serverRoute)-1] != '/' {
		r.Get(serverRoute, http.RedirectHandler(serverRoute+"/", http.StatusMovedPermanently).ServeHTTP)
		serverRoute += "/"
	}
	serverRoute += "*"

	r.Get(serverRoute, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		serverRoutePrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(serverRoutePrefix, http.FileServer(pathToStaticFolder))
		fs.ServeHTTP(w, r)
	})
}
