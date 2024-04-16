package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gerry-sheva/greenlight.git/internal/data"
	"github.com/go-chi/chi/v5"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	movie := data.Movies{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"War", "Romance"},
		Version:   1,
	}
	err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
