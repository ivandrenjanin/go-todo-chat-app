package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	pg "github.com/ivandrenjanin/go-chat-app/db/pg/pg_generated"
	"github.com/ivandrenjanin/go-chat-app/pkg/cfg"
)

type Database struct {
	Pg *pg.Queries
	Db *sql.DB
	// TODO: mongo type whatever that will be
}

func New(config *cfg.Config) (Database, error) {
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
		Pg: pg.New(pgConn),
		Db: pgConn,
	}

	return database, nil
}
