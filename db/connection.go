package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/ivandrenjanin/go-chat-app/cfg"
)

type Database struct {
	Queries *Queries
	// TODO: mongo type whatever that will be
}

func CreateDBConn(config *cfg.Config) (Database, error) {
	pgConnStr := fmt.Sprintf(
		"dbname=%s user=%s password=%s sslmode=%s host=%s port=%d",
		config.PgConfig.DbName,
		config.PgConfig.User,
		config.PgConfig.Password,
		config.PgConfig.SslMode,
		config.PgConfig.Host,
		config.PgConfig.Port,
	)
	pgConn, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		return Database{}, fmt.Errorf("db connection err: %s\n", err)
	}

	database := Database{
		Queries: New(pgConn),
	}

	return database, nil
}
