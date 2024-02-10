package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ivandrenjanin/go-chat-app/cfg"
	"github.com/ivandrenjanin/go-chat-app/db"
)

func CreateServer(config *cfg.Config) error {
	db, err := db.CreateDBConn(config)
	if err != nil {
		return err
	}

	mux := chi.NewRouter()
	addRoutes(mux, config, &db)

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
