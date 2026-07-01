package service

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(context.Context, string, string) error
	LoginUser(context.Context, string, string) (string, string, error)
}

type TaskService interface {
	Create(context.Context, *model.Task) error
	FindByID(context.Context, uuid.UUID, uuid.UUID) (*model.Task, error)
	FindAllByUserID(context.Context, uuid.UUID) ([]*model.Task, error)
	Update(context.Context, *model.Task) error
	Delete(context.Context, uuid.UUID, uuid.UUID) error
}

type Models struct {
	UserService
	TaskService
}

func NewServiceModels(RepoModels *repository.Models) *Models {
	return &Models{
		NewUserService(RepoModels.UserRepository),
		NewTaskService(RepoModels.TaskRepository),
	}
}
