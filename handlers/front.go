package handlers

import (
	"canvas/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FrontPage(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = render(w, r, views.FrontPage())
	})
}
