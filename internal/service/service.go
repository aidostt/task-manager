package service

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/aidostt/task-manager/pkg/jwt"
	"github.com/google/uuid"
)

type User interface {
	RegisterUser(context.Context, string, string) (string, string, error)
	LoginUser(context.Context, string, string) (string, string, error)
}

type Task interface {
	Create(context.Context, *model.Task) error
	FindByID(context.Context, uuid.UUID, uuid.UUID) (*model.Task, error)
	FindAllByUserID(context.Context, uuid.UUID) ([]*model.Task, error)
	Update(context.Context, *model.Task) error
	Delete(context.Context, uuid.UUID, uuid.UUID) error
}

type Models struct {
	User
	Task
}

func NewServiceModels(RepoModels *repository.Models, manager jwt.TokenManager) *Models {
	return &Models{
		NewUserService(RepoModels.UserRepository, manager, RepoModels.SessionRepository),
		NewTaskService(RepoModels.TaskRepository),
	}
}
