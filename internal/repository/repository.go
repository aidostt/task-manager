package repository

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(context.Context, *model.User) error
	GetUserByEmail(context.Context, string) (*model.User, error)
	GetUserByID(context.Context, uuid.UUID) (*model.User, error)
	UpdateUser(context.Context, *model.User) error
	DeleteUser(context.Context, uuid.UUID) error
}

type TaskRepository interface {
	CreateTask(context.Context, *model.Task) error
	GetTaskByID(context.Context, uuid.UUID) (*model.Task, error)
	GetTasksByUserID(context.Context, uuid.UUID) ([]*model.Task, error)
	UpdateTask(context.Context, *model.Task) error
	DeleteTask(context.Context, uuid.UUID) error
}
