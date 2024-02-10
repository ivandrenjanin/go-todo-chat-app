package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/api/handlers"
	"github.com/ivandrenjanin/go-chat-app/services"
)

func addRoutes(
	mux *chi.Mux,
	us *services.UserService,
	ps *services.ProjectService,
	ts *services.ToDoService,
) {
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer(mux, "/files", filesDir)

	mux.Route("/", func(r chi.Router) {
		r.Get("/", handlers.PublicHomeHandler())
	})

	mux.Route("/api/users", func(r chi.Router) {
		// Public Routes
		// Login
		// Register

		// Everything else is protected
	})

	mux.Route("/api/projects", func(r chi.Router) {
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})
}
