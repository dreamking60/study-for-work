package main

import (
	"log"

	"edge-gateway/internal/server"
)

func main() {
	srv := server.New()
	log.Printf("edge-gateway starting on %s", srv.Addr())

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
