package main

import (
	"log"

	"CRUD/internal/config"
	"CRUD/internal/db"
	"CRUD/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	client, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	s := server.NewServer(cfg, client)
	s.Run()
}
