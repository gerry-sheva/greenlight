package main

import "github.com/go-chi/chi/v5"

func (app *appliction) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/v1/healthcheck", app.healthCheckHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	return router
}
