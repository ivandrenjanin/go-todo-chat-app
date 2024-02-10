package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/cfg"
	"github.com/ivandrenjanin/go-chat-app/db"
)

func addRoutes(
	mux *chi.Mux,
	config *cfg.Config,
	db *db.Database,
) {
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeJSON))

	mux.Route("/api/users", func(r chi.Router) {
	})

	mux.Route("/api/projects", func(r chi.Router) {
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})
}
