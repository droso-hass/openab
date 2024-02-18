package utils

import (
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/render"
)

type response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func IfError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err == nil {
		return false
	}
	Error(w, r, 500, err.Error())
	return true
}

func Error(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	resp := response{Status: "error", Data: msg}
	render.JSON(w, r, resp)
}

func JSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	render.Status(r, code)
	render.JSON(w, r, response{Status: "ok", Data: payload})
}

func SendFile(w http.ResponseWriter, r *http.Request, path string, contentType string) {
	w.Header().Set("Content-Type", contentType)
	data, err := os.ReadFile(path)
	if err != nil {
		render.Status(r, 404)
		slog.Error("error reading file", "path", path)
	}
	w.Write(data)
}

func GetIPFromRequest(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	h, _, err := net.SplitHostPort(ip)
	if err == nil {
		return h
	}
	return ip
}
