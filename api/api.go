package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
	"github.com/ivandrenjanin/go-chat-app/pkg/cfg"
	"github.com/ivandrenjanin/go-chat-app/pkg/mailer"
	projectStore "github.com/ivandrenjanin/go-chat-app/store/project"
	todoStore "github.com/ivandrenjanin/go-chat-app/store/todo"
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
	todoStore := todoStore.New(&db)

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
	todoService := app.NewTodoService(&todoStore)
	projectService := app.NewProjectService(&projectStore, &mailer)

	mux := chi.NewRouter()
	err = addRoutes(mux, &userService, &projectService, &todoService, &identityService)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf(
		"%s:%d",
		config.AppConfig.Host,
		config.AppConfig.Port,
	)

	srv := &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	log.Printf("Starting a server on %s 🚀\n", addr)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
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
