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
		ch.ServeHTTP(w, r)
	}
}
