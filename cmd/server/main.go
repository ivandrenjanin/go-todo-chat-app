package main

import (
	"log"

	"github.com/ivandrenjanin/go-chat-app/api"
)

func main() {
	srv := api.CreateServer()
	log.Fatal(srv.ListenAndServe())
}
