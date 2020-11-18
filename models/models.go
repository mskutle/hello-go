package models

import (
	"time"
)

type User struct {
	Name     string `json: "name" validate: "required"`
	Username string `json: "username" validate: "required"`
	Password string `json: "password" validate: "required"`
}

type Todo struct {
	Title       string    `json: "title" validate: "required"`
	CreatedAt   time.Time `json: "createdAt validate: "required"`
	CompletedAt time.Time `json: "completedAt validate: "required"`
}

type ErrorResponse struct {
	Message string
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}

func NewUser(name, username, password string) User {
	return User{Username: username, Password: password, Name: name}
}

func NewTodo(title string) Todo {
	return Todo{Title: title, CreatedAt: time.Now()}
}
