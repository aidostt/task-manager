package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.PasswordHash)
	return err
}

func (r *UserRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	user := &model.User{}
	err := r.db.GetContext(ctx, user, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return user, err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	user := &model.User{}
	err := r.db.GetContext(ctx, user, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return user, err
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET email = $1, password_hash = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.PasswordHash, user.ID)
	return err
}

func (r *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
