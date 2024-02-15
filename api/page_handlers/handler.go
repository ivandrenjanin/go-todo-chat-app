package pagehandler

import (
	"net/http"

	"github.com/a-h/templ"

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
func IndexPageProtected() http.HandlerFunc {
	ch := templ.Handler(pages.IndexProtected())
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle this in a middleware + Validate JWT
		_, err := r.Cookie("app-token")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ch.ServeHTTP(w, r)
	}
}
