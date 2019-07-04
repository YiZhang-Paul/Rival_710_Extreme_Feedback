package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yi-zhang/rival-710-extreme-feedback/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	router = chi.NewRouter()
)

func init() {
	// middlewares
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	// default response for http methods (405 if not specified)
	router.Post("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
}

func getRoutes() *chi.Mux {
	router.Post("/notify", func(w http.ResponseWriter, r *http.Request) {
		body, ok := utils.ParseJSON(r.Body)
		if ok {
			log.Print(body)
		}
	})
	return router
}
