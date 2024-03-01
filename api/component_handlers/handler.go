package componenthandlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

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

func EditProjectComponent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
