package userRoute

import (
	user_handlers "CRUD/handlers"

	"github.com/go-chi/chi/v5"
)

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/user/getUsers", user_handlers.GetUsers)
	r.Post("/api/user/createUser", user_handlers.SaveUser)
	r.Put("/api/user/updateUser", user_handlers.UpdateUser)
	r.Delete("/api/user/deleteUser", user_handlers.DeleteUser)

	return r
}
