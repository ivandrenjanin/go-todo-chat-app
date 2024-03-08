package cmd

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"

	"github.com/ivandrenjanin/go-chat-app/api"
	"github.com/ivandrenjanin/go-chat-app/pkg/cfg"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves api over http",
	Run:   runServe,
}

func runServe(cmd *cobra.Command, args []string) {
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

	if err := api.New(&config); err != nil {
		log.Fatalln(err)
	}
}
