package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"

	"github.com/ivandrenjanin/go-chat-app/cfg"
)

// TODO: Use flags here to run migration up
var loadEnv = flag.Bool("load-env", false, "load local .env file")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if *loadEnv {
		err := godotenv.Load()
		if err != nil {
			return err
		}

	}

	config, err := cfg.CreateConfig()
	if err != nil {
		return err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.PgConfig.User,
		config.PgConfig.Password,
		config.PgConfig.Host,
		config.PgConfig.Port,
		config.PgConfig.DbName,
		config.PgConfig.SslMode,
	)

	m, err := migrate.New(
		"file://migrations",
		connStr,
	)
	if err != nil {
		log.Fatalf("migration failed: %s", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("migration:up failed: %s", err)
	}

	return nil
}
