package repository

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(ctx context.Context, task *model.Task) error {
	query := `INSERT INTO tasks (user_id, title, description, status, priority) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, task.UserID, task.Title, task.Description, task.Status, task.Priority)
	return err
}

func (r *taskRepo) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	query := `SELECT * FROM tasks WHERE id = $1`
	task := &model.Task{}
	err := r.db.GetContext(ctx, task, query, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepo) GetTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Task, error) {
	query := `SELECT * FROM tasks WHERE user_id = $1`
	tasks := []*model.Task{}
	err := r.db.SelectContext(ctx, &tasks, query, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *model.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, priority = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.Priority, task.ID)
	return err
}

func (r *taskRepo) DeleteTask(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
