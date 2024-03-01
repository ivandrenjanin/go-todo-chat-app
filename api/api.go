package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
	"github.com/ivandrenjanin/go-chat-app/pkg/cfg"
	"github.com/ivandrenjanin/go-chat-app/pkg/mailer"
	projectStore "github.com/ivandrenjanin/go-chat-app/store/project"
	userStore "github.com/ivandrenjanin/go-chat-app/store/user"
)

func New(config *cfg.Config) error {
	db, err := db.New(config)
	if err != nil {
		return err
	}

	// Define Stores
	userStore := userStore.New(&db)
	projectStore := projectStore.New(&db)

	// Define Packages
	mailer := mailer.New(
		config.MailerConfig.Host,
		config.MailerConfig.Username,
		config.MailerConfig.Password,
		config.MailerConfig.Port,
	)

	// Define Services
	identityService := app.NewIdentityService(&config.JwtConfig, &userStore)
	userService := app.NewUserService(&userStore)
	todoService := app.NewTodoService()
	projectService := app.NewProjectService(&projectStore, &mailer)

	mux := chi.NewRouter()
	err = addRoutes(mux, &userService, &projectService, &todoService, &identityService)
	if err != nil {
		return err
	}

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
