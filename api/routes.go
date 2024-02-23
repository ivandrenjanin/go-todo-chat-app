package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	ah "github.com/ivandrenjanin/go-chat-app/api/api_handlers"
	ch "github.com/ivandrenjanin/go-chat-app/api/component_handlers"
	ph "github.com/ivandrenjanin/go-chat-app/api/page_handlers"
	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/components"
)

func addRoutes(
	mux *chi.Mux,
	us *app.UserService,
	ps *app.ProjectService,
	ts *app.ToDoService,
	is *app.IdentityService,
) error {
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer(mux, "/files", filesDir)

	// Handle Pages
	mux.Route("/", func(r chi.Router) {
		// Public Pages
		r.Get("/", ph.IndexPage())

		// Protected Pages
		r.Group(func(r chi.Router) {
			r.Use(MakeIdentityMiddleware(is, us))
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
		// r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		// 	ruid := chi.URLParam(r, "userId")
		// 	uid, err := strconv.Atoi(ruid)
		// 	if err != nil {
		// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		// 	}
		//
		// 	p, err := ps.FindProjectsByUserId(r.Context(), uid)
		// 	fmt.Printf("Projects: %v\n", p)
		// })

		r.Use(MakeIdentityMiddleware(is, us))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			u := r.Context().Value("user").(app.User)
			projects, _ := ps.FindProjectsByUserId(r.Context(), u.ID)

			headers := []string{"Project Name", "Description", "Actions"}
			rows := make([][]string, 0, cap(projects))
			base := "/api/projects"

			for _, project := range projects {
				subBase := fmt.Sprintf("%s/%d", base, project.ID)
				assign := fmt.Sprintf("%s/assign", subBase)
				r := []string{assign, subBase, project.Name, project.Description}
				rows = append(rows, r)
			}

			c := templ.Handler(components.ProjectTable(headers, rows))
			c.ServeHTTP(w, r)
		})
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})

	return nil
}
