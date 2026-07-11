package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      string    `db:"status" json:"status"`
	Priority    string    `db:"priority" json:"priority"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

type TaskIdInput struct {
	ID string `db:"id" json:"id"`
}
