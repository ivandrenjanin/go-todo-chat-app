package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	ah "github.com/ivandrenjanin/go-chat-app/api/api_handlers"
	ch "github.com/ivandrenjanin/go-chat-app/api/component_handlers"
	ph "github.com/ivandrenjanin/go-chat-app/api/page_handlers"
	"github.com/ivandrenjanin/go-chat-app/app"
)

type session struct {
	email  string
	expiry time.Time
}

type sessionMap = map[string]session

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func addRoutes(
	mux *chi.Mux,
	us *app.UserService,
	ps *app.ProjectService,
	ts *app.ToDoService,
	is *app.IdentityService,
) {
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer(mux, "/files", filesDir)

	// Handle Pages
	mux.Route("/", func(r chi.Router) {
		// Public Pages
		r.Get("/", ph.IndexPage())

		// Protected Pages
		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					c, err := r.Cookie("session_token")
					if err != nil {
						http.Error(
							w,
							http.StatusText(http.StatusUnauthorized),
							http.StatusUnauthorized,
						)
						fmt.Println("No cookie found!")
						return
					}

					token := c.Value
					claims, ok := is.ValidateToken(token)
					if !ok {
						http.Error(
							w,
							http.StatusText(http.StatusUnauthorized),
							http.StatusUnauthorized,
						)
						fmt.Println("Invalid token")
						return
					}

					u, err := us.FindById(r.Context(), claims.UserID)
					if err != nil {
						http.Error(
							w,
							http.StatusText(http.StatusUnauthorized),
							http.StatusUnauthorized,
						)
						fmt.Println("User not found")
						return

					}

					ctx := context.WithValue(r.Context(), "user", u)
					next.ServeHTTP(w, r.WithContext(ctx))
				})
			})
			r.Get("/home", ph.IndexPageProtected(us, ps))
		})
	})

	// Handle Components
	mux.Route("/api/components", func(r chi.Router) {
		// Public Components
		r.Get("/home-page-form/", ch.HomePageFormComponent())

		// Protected Components
	})

	// Handle Api
	mux.Route("/api/auth", func(r chi.Router) {
		// Public Routes
		r.Post("/register", ah.RegisterHandler(is))

		r.Post("/login", ah.LoginHandler(is))
	})

	mux.Route("/api/users", func(r chi.Router) {
	})

	mux.Route("/api/projects", func(r chi.Router) {
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})
}
