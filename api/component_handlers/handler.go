package componenthandlers

import (
	"net/http"

	"github.com/a-h/templ"

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
