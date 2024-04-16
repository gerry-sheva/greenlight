package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *appliction) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	router.Get("/v1/healthcheck", app.healthCheckHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	return router
}
