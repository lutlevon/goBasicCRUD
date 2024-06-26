package user_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "Von Guzman"},
	{ID: "2", Name: "Von Sai"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, newUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
