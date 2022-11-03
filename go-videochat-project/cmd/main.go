package main

import (
	"log"

	"github.com/worachanon226/go-videochat-project/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
