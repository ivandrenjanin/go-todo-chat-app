package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"

	"github.com/ivandrenjanin/go-chat-app/api"
	"github.com/ivandrenjanin/go-chat-app/cfg"
)

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

	ctx := context.Background()
	var config cfg.Config

	err := envconfig.Process(ctx, &config)
	if err != nil {
		return err
	}

	err = api.New(&config)

	return err
}
