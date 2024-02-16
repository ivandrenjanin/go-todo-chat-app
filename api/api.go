package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/cfg"
	"github.com/ivandrenjanin/go-chat-app/db"
	"github.com/ivandrenjanin/go-chat-app/storage"
)

func New(config *cfg.Config) error {
	db, err := db.New(config)
	if err != nil {
		return err
	}

	mux := chi.NewRouter()
	userStorage := storage.NewUserStorage(&db)

	as := app.NewAuthService(&config.JwtConfig, &userStorage)
	us := app.NewUserService()
	ts := app.NewTodoService()
	ps := app.NewProjectService()

	addRoutes(mux, &us, &ps, &ts, &as)

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
