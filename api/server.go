package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/api/handlers"
	"github.com/ivandrenjanin/go-chat-app/cfg"
	"github.com/ivandrenjanin/go-chat-app/db"
)

func CreateServer(config *cfg.Config) *http.Server {
	db, err := db.CreateDBConn(config)
	if err != nil {
		log.Fatalf("can not connect to db: %s", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	handlers.CreateUserHandlers(router, db)

	srv := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(
			"%s:%d",
			config.AppConfig.Host,
			config.AppConfig.Port,
		),
	}

	return srv
}
