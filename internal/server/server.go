package server

import (
	"CRUD/ent"
	"CRUD/internal/config"
	"CRUD/internal/handlers"
	"CRUD/internal/routes"
	"CRUD/internal/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    *config.Config
	client *ent.Client
	router *chi.Mux
}

func NewServer(cfg *config.Config, client *ent.Client) *Server {
	userService := services.NewUserService(client)
	userHandler := handlers.NewUserHandler(userService)

	router := routes.NewRouter(userHandler)

	return &Server{
		cfg:    cfg,
		client: client,
		router: router,
	}
}

func (s *Server) Run() {
	log.Println("Server is runing on port 3000")
	http.ListenAndServe(":3000", s.router)
}
