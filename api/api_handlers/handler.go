package apihandlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"

	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body %s", err)
			return
		}

		fmt.Printf("Body: %s\n", string(body))
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body %s", err)
			return
		}

		fmt.Printf("Body: %s\n", string(body))
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}
