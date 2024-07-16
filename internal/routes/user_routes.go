package routes

import (
	"CRUD/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(userHandler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", userHandler.GetUsers)
	r.Get("/{id}", userHandler.GetUserById)
	return r
}
