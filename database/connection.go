package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	Queries *Queries
	// mongo type whatever that will be
}

func CreateDBConn() (Database, error) {
	// TODO: Load these from an env
	connStr := "dbname=postgres user=postgres password=postgres sslmode=disable host=localhost port=5432"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return Database{}, fmt.Errorf("db connection err: %s\n", err)
	}

	database := Database{
		Queries: New(conn),
	}

	return database, nil
}
