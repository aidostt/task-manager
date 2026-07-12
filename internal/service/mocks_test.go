package service

import (
	"context"
	"time"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/pkg/jwt"
	"github.com/google/uuid"
)

type mockUserRepo struct {
	getUser   *model.User
	getErr    error
	createErr error
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *model.User) (uuid.UUID, error) {
	return uuid.New(), m.createErr
}
func (m *mockUserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.getUser, m.getErr
}
func (m *mockUserRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return m.getUser, m.getErr
}
func (m *mockUserRepo) UpdateUser(ctx context.Context, user *model.User) error {
	return m.createErr
}
func (m *mockUserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return m.createErr
}

// мок SessionRepository
type mockSessionRepo struct {
	session *model.Session
	err     error
}

func (m *mockSessionRepo) GetByToken(ctx context.Context, token string) (*model.Session, error) {
	return m.session, m.err
}
func (m *mockSessionRepo) CreateSession(ctx context.Context, session *model.Session) error {
	return m.err
}
func (m *mockSessionRepo) DeleteSession(ctx context.Context, token string) error {
	return m.err
}

// мок TokenManager
type mockTokenManager struct{}

func (m *mockTokenManager) NewAccessToken(userID string) (string, error) {
	return "access_token", nil
}
func (m *mockTokenManager) NewRefreshToken() (string, error) { return "refresh_token", nil }
func (m *mockTokenManager) Parse(token string) (*jwt.Claims, error) {
	return &jwt.Claims{UserID: token}, nil
}
func (m *mockTokenManager) RefreshTTL() time.Duration { return time.Hour }
