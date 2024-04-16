package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Show this movie: %s\n", id)
}
