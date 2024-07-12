package main

import (
	"log"
	"net/http"

	"CRUD/config"
	"CRUD/db"
	userRoute "CRUD/routes"
)

func main() {
	cfg := config.LoadConfig()

	client, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	r := userRoute.GetRoutes(client)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("could not start server")
	}
}
