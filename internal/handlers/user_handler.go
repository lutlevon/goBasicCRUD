package handlers

import (
	"CRUD/internal/services"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers(context.Background())
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	u, err := h.userService.GetUserById(context.Background(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

// func GetUserById(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	for _, user := range users {
// 		if user.ID == id {
// 			w.Header().Set("content-Type", "application/json")
// 			json.NewEncoder(w).Encode(user)
// 			return
// 		}
// 	}
// 	http.Error(w, "User not found", http.StatusNotFound)
// }

// func SaveUser(w http.ResponseWriter, r *http.Request) {
// 	var newUser User
// 	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	users = append(users, newUser)
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("content-Type", "application/json")
// 	json.NewEncoder(w).Encode(newUser)
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	var updatedUser User
// 	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	for i, user := range users {
// 		if user.ID == id {
// 			users[i] = updatedUser
// 			w.Header().Set("content-Type", "application/json")
// 			json.NewEncoder(w).Encode(users[i])
// 			return
// 		}
// 	}

// 	http.Error(w, "User not found", http.StatusNotFound)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")

// 	for i, user := range users {
// 		if user.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			w.Header().Set("content-Type", "application/json")
// 			json.NewEncoder(w).Encode(users)
// 			return
// 		}
// 	}

// 	http.Error(w, "User not found", http.StatusNotFound)
// }
