package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/api/handlers"
	"github.com/ivandrenjanin/go-chat-app/services"
	"github.com/ivandrenjanin/go-chat-app/views/pages"
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
		r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
			templ.Handler(pages.IndexPrivate()).ServeHTTP(w, r)
		})
	})

	mux.Route("/api/users", func(r chi.Router) {
		// Public Routes
		// Login
		// Register
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading the body %s", err)
				return
			}

			fmt.Printf("Body: %s\n", string(body))
			w.Header().Add("HX-Push-Url", "home")
			ch := templ.Handler(pages.IndexPrivate())
			ch.ServeHTTP(w, r)
		})
		// Everything else is protected
	})

	mux.Route("/api/projects", func(r chi.Router) {
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})
}
