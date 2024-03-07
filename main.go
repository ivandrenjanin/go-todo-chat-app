package main

import (
	"log"

	"github.com/ivandrenjanin/go-chat-app/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
