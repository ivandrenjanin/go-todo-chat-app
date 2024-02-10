package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// TODO: Use flags here to create new migration files
// TODO: Use flags here to run migration up
func main() {
	m, err := migrate.New(
		"file://migrations",
		// TODO: import from config env
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("migration failed: %s", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("migration:up failed: %s", err)
	}
}
