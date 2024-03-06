package componenthandlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/components"
)

func HomePageFormComponent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		if len(q) <= 0 {
			return
		}

		if q == "login" {
			templ.Handler(components.HomePageForm("signup", components.SignupFormFields)).
				ServeHTTP(w, r)
			return
		}

		if q == "signup" {
			templ.Handler(components.HomePageForm("login", components.LoginFormFields)).
				ServeHTTP(w, r)
			return
		}

		return
	}
}

func AssignUserToProjectComponent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubId := chi.URLParam(r, "projectId")

		url := fmt.Sprintf("/api/projects/%s/invitations", pubId)

		ch := templ.Handler(components.AssignUserModal(url))
		ch.ServeHTTP(w, r)
	}
}

func EditProjectComponent(ps *app.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubId := chi.URLParam(r, "projectId")
		p, err := ps.FindProjectById(r.Context(), pubId)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		u := r.Context().Value("user").(app.User)
		if p.OwnerID != u.ID {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		var m map[string]string = map[string]string{
			"name":        p.Name,
			"description": p.Description,
		}

		url := fmt.Sprintf("/api/projects/%s", pubId)

		ch := templ.Handler(components.EditProjectModal(url, m))
		ch.ServeHTTP(w, r)
	}
}
