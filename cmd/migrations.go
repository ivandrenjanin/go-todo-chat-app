package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"

	"github.com/ivandrenjanin/go-chat-app/pkg/cfg"
)

func init() {
	rootCmd.AddCommand(migrateUpCmd)
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "runs migrate up command",
	Long:  "runs migrate up command",
	Run:   runMigrateUp,
}

func runMigrateUp(cmd *cobra.Command, args []string) {
	if loadCfg {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}

	}

	ctx := context.Background()
	var config cfg.Config

	err := envconfig.Process(ctx, &config)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	if err := m.Up(); err != nil {
		log.Fatalln(err)
		// TODO: Look into this
		// err = m.Down()
		// if err != nil {
		// 	return err
		// }
	}
}
