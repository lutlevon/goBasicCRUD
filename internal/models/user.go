package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
