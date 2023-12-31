package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"datareporter/config"
	"datareporter/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var env = config.AppConfig{}
var cArticle = controller.Article{}

func main() {
	config, err := env.Load()
	if err != nil {
		log.Println(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/articles", cArticle.List)

	log.Printf("Starting up on http://%s:%s", config.Host, config.Port)

	http.ListenAndServe(":"+config.Port, r)
}
