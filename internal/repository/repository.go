package repository

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(context.Context, *model.User) (uuid.UUID, error)
	GetUserByEmail(context.Context, string) (*model.User, error)
	GetUserByID(context.Context, uuid.UUID) (*model.User, error)
	UpdateUser(context.Context, *model.User) error
	DeleteUser(context.Context, uuid.UUID) error
}

type TaskRepository interface {
	CreateTask(context.Context, *model.Task) (*model.Task, error)
	GetTaskByID(context.Context, uuid.UUID) (*model.Task, error)
	GetTasksByUserID(context.Context, uuid.UUID) ([]*model.Task, error)
	UpdateTask(context.Context, *model.Task) error
	DeleteTask(context.Context, uuid.UUID) error
}

type SessionRepository interface {
	CreateSession(context.Context, *model.Session) error
	GetByToken(context.Context, string) (*model.Session, error)
	DeleteSession(context.Context, string) error
}

type Models struct {
	UserRepository
	TaskRepository
	SessionRepository
}

func NewRepoModels(db *sqlx.DB) *Models {
	return &Models{
		NewUserRepo(db),
		NewTaskRepo(db),
		NewSessionRepo(db),
	}
}
