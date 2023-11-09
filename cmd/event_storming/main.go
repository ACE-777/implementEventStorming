package main

import (
	"log"

	server "implementEventStorming/internal/server"
)

func main() {
	log.Printf("start eventStorming model")

	server.RunServer()
}
