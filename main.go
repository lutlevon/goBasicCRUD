package main

import (
	"log"
	"net/http"

	userRoute "CRUD/routes"
)

func main() {
	r := userRoute.GetRoutes()

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("could not start server")
	}
}
