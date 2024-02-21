package pagehandler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

// Public Pages
func IndexPage() http.HandlerFunc {
	ch := templ.Handler(pages.Index())

	return func(w http.ResponseWriter, r *http.Request) {
		ch.ServeHTTP(w, r)
	}
}

// Protected Pages
func IndexPageProtected(us *app.UserService, ps *app.ProjectService) http.HandlerFunc {
	ch := templ.Handler(pages.IndexProtected())
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value("user").(app.User)
		fmt.Printf("User: %#v\n", u)
		ch.ServeHTTP(w, r)
	}
}
