package service

import (
	"context"
	"errors"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/google/uuid"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(ctx context.Context, task *model.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	return t.repo.CreateTask(ctx, task)
}

func (t *TaskService) FindByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.Task, error) {
	task, err := t.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, errors.New("forbidden")
	}
	return task, nil
}

func (t *TaskService) FindAllByUserID(ctx context.Context, id uuid.UUID) ([]*model.Task, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid task id")
	}
	return t.repo.GetTasksByUserID(ctx, id)
}

func (t *TaskService) Update(ctx context.Context, task *model.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	taskFromBD, err := t.repo.GetTaskByID(ctx, task.ID)
	if err != nil {
		return err
	}
	if taskFromBD.UserID != task.UserID {
		return errors.New("forbidden")
	}
	return t.repo.UpdateTask(ctx, task)
}

func (t *TaskService) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid task id")
	}
	task, err := t.repo.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("forbidden")
	}
	return t.repo.DeleteTask(ctx, id)
}
