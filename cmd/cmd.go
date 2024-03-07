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

var (
	loadCfg bool
	rootCmd = &cobra.Command{
		Use:   "serve",
		Short: "serves the http server",
		Long:  "serves the http server",
		Run:   run,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().
		BoolVar(&loadCfg, "load-env", false, "load-env in order to load local .env file")
}

func run(cmd *cobra.Command, args []string) {
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
