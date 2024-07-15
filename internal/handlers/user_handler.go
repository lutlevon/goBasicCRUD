package handlers

import (
	"CRUD/ent"
	"context"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.client.User.Query().All(context.Background())
	if err != nil {
		http.Error(w, "failed to query users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
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
