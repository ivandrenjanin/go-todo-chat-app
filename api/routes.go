package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	ah "github.com/ivandrenjanin/go-chat-app/api/api_handlers"
	ch "github.com/ivandrenjanin/go-chat-app/api/component_handlers"
	ph "github.com/ivandrenjanin/go-chat-app/api/page_handlers"
	"github.com/ivandrenjanin/go-chat-app/app"
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
			r.Get("/project/{projectId}", ph.ProjectPageProtected(ps, ts))
		})
	})

	// Handle Components
	mux.Route("/api/public/components", func(r chi.Router) {
		// Public Components
		r.Get("/home-page-form/", ch.HomePageFormComponent())
	})

	mux.Route("/api/p", func(r chi.Router) {
		r.Get("/accept-invitation/", ah.AcceptInvitationHandler(us, ps))
	})

	mux.Route("/api/components", func(r chi.Router) {
		// Protected Components
		r.Use(MakeIdentityMiddleware(is, us))
		r.Get("/assign-user-project/{projectId}", ch.AssignUserToProjectComponent())
		r.Get("/edit-project/{projectId}", ch.EditProjectComponent(ps))
	})

	// Handle Api
	mux.Route("/api/auth", func(r chi.Router) {
		// Public Routes
		r.Post("/register", ah.RegisterHandler(is))
		r.Post("/login", ah.LoginHandler(is))
	})

	mux.Route("/api/projects", func(r chi.Router) {
		// Protected Routes
		r.Use(MakeIdentityMiddleware(is, us))
		r.Get("/", ah.GetProjectByUserIdHandler(ps))
		r.Delete("/{projectId}", ah.DeleteProjectHandler(ps))
		r.Patch("/{projectId}", ah.EditProjectHandler(ps))
		r.Post("/", ah.CreateProjectHandler(ps))
		r.Post("/{projectId}/invitations", ah.CreateProjectInvitationHandler(ps))
	})

	mux.Route("/api/users", func(r chi.Router) {
	})

	mux.Route("/api/todo", func(r chi.Router) {
	})

	return nil
}
