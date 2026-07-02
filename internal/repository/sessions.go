package repository

import (
	"context"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/jmoiron/sqlx"
)

type SessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) CreateSession(ctx context.Context, session *model.Session) error {
	query := `INSERT INTO sessions (user_id, token, expires_at) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, session.UserID, session.Token, session.ExpiresAt)
	return err
}

func (r *SessionRepo) GetByToken(ctx context.Context, token string) (*model.Session, error) {
	query := `SELECT * FROM sessions WHERE token = $1`
	session := &model.Session{}
	err := r.db.GetContext(ctx, session, query, token)
	if err != nil {
		return nil, err
	}
	return session, nil
}
func (r *SessionRepo) DeleteSession(ctx context.Context, token string) error {
	query := `DELETE FROM sessions WHERE token = $1`
	_, err := r.db.ExecContext(ctx, query, token)
	return err
}
