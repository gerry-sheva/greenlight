package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.notFoundErrorResponse(w, r)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		app.methodNotAllowedErrorResponse(w, r)
	})

	router.Get("/v1/healthcheck", app.healthCheckHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	return router
}
