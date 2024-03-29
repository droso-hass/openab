package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/droso-hass/openab/internal/api"
	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/config"
	"github.com/droso-hass/openab/internal/udp"
	"github.com/droso-hass/openab/internal/utils"
	v2 "github.com/droso-hass/openab/internal/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	config.Parse("./config.yaml")
	utils.SetupLogs("debug")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	err := udp.Start(4000)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New("nats://127.0.0.1:4222")

	tagtag := v2.New(r, api)

	api.Listen([]common.INabReceiverHander{tagtag})

	utils.ServeStatic(r, "/data", http.Dir("./static"))

	fmt.Println("Ready !")
	err = http.ListenAndServe(":80", r)
	log.Fatal(err)
}
