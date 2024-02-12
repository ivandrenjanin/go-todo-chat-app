package componenthandlers

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

func HomePageFormComponent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
