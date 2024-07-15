package routes

import (
	"CRUD/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *handlers.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/users", UserRoutes(userHandler))

	return r

}
