package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoRequest struct {
	PreviousTitle *string `json:"previous_title,omitempty"`
	UpdatedTitle  *string `json:"updated_title,omitempty"`
	IsCompleted   *bool   `json:"is_completed,omitempty"`
}

type TodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
