package service

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/google/uuid"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) Task {
	return &taskService{repo: repo}
}

func (t *taskService) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	if task == nil {
		return nil, ErrInvalidTask
	}
	return t.repo.CreateTask(ctx, task)
}

func (t *taskService) FindByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.Task, error) {
	task, err := t.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, ErrForbidden
	}
	return task, nil
}

func (t *taskService) FindAllByUserID(ctx context.Context, id uuid.UUID) ([]*model.Task, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidUserID
	}
	return t.repo.GetTasksByUserID(ctx, id)
}

func (t *taskService) Update(ctx context.Context, task *model.Task) error {
	if task == nil {
		return ErrInvalidTask
	}
	taskFromBD, err := t.repo.GetTaskByID(ctx, task.ID)
	if err != nil {
		return err
	}
	if taskFromBD.UserID != task.UserID {
		return ErrForbidden
	}
	return t.repo.UpdateTask(ctx, task)
}

func (t *taskService) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidTaskID
	}
	task, err := t.repo.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return ErrForbidden
	}
	return t.repo.DeleteTask(ctx, id)
}
