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

	// 3 Types of handlers
	// A. Entire Page Handler
	//    1a - Public pages
	//    2a - Protected pages
	// B. Api handlers
	//    1b - Public Api
	//    2b - Protected Api
	// C. Component handlers
	//    1c - Public components
	//    2c - Protected components

	mux.Route("/", func(r chi.Router) {
		r.Get("/", templ.Handler(pages.Index()).ServeHTTP)

		r.Get("/home", templ.Handler(pages.IndexPrivate()).ServeHTTP)
	})

	mux.Route("/api/components", func(r chi.Router) {
		r.Get("/home-page-form/", func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query().Get("q")

			if len(q) <= 0 {
				return
			}

			if q == "login" {
				templ.Handler(pages.HomePageForm("signup", pages.SignupFormFields)).ServeHTTP(w, r)
				return
			}

			if q == "signup" {
				templ.Handler(pages.HomePageForm("login", pages.LoginFormFields)).ServeHTTP(w, r)
				return
			}

			return
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

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
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
