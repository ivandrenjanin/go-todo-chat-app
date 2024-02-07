package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/api/handlers"
	"github.com/ivandrenjanin/go-chat-app/database"
)

func CreateServer() *http.Server {
	db, err := database.CreateDBConn()
	if err != nil {
		log.Fatalf("can not connect to db: %s", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	handlers.CreateUserHandlers(router, db)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
	}

	return srv
}
