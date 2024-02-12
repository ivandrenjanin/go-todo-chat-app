package handlers

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

func PublicHomeHandler() http.HandlerFunc {
	// we do the thing
	return templ.Handler(pages.Index()).ServeHTTP
}
