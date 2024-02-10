package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/cfg"
	"github.com/ivandrenjanin/go-chat-app/db"
	"github.com/ivandrenjanin/go-chat-app/services"
)

func CreateServer(config *cfg.Config) error {
	db, err := db.CreateDBConn(config)
	if err != nil {
		return err
	}

	mux := chi.NewRouter()
	us := services.NewUser(&db)
	ts := services.NewTodo(&db)
	ps := services.NewProject(&db)

	addRoutes(mux, &us, &ps, &ts)

	srv := &http.Server{
		Handler: mux,
		Addr: fmt.Sprintf(
			"%s:%d",
			config.AppConfig.Host,
			config.AppConfig.Port,
		),
	}

	err = srv.ListenAndServe()

	return err
}

func fileServer(mux *chi.Mux, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		mux.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	mux.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
