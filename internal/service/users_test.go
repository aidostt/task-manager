package service

import (
	"context"
	"testing"
	"time"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	svc := NewUserService(
		&mockUserRepo{getUser: &model.User{Email: "test@test.com"}, createErr: nil},
		&mockTokenManager{},
		&mockSessionRepo{},
	)

	_, _, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")

	assert.ErrorIs(t, err, ErrUserAlreadyExists)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	svc := NewUserService(
		&mockUserRepo{getUser: &model.User{PasswordHash: string(hash)}, createErr: nil},
		&mockTokenManager{},
		&mockSessionRepo{},
	)

	_, _, err := svc.LoginUser(context.Background(), "test@test.com", "password123")
	assert.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestRefreshSession_SessionExpired(t *testing.T) {
	svc := NewUserService(
		&mockUserRepo{},
		&mockTokenManager{},
		&mockSessionRepo{session: &model.Session{
			UserID:    uuid.New(),
			ExpiresAt: time.Now().Add(-time.Hour), // час назад
		}},
	)

	_, _, err := svc.RefreshTokens(context.Background(), "")
	assert.ErrorIs(t, err, ErrSessionExpired)
}

func TestRegisterUser_Success(t *testing.T) {
	svc := NewUserService(
		&mockUserRepo{getUser: nil, createErr: nil},
		&mockTokenManager{},
		&mockSessionRepo{},
	)

	accessToken, refreshToken, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "access_token", accessToken)
	assert.Equal(t, "refresh_token", refreshToken)
}

func TestLoginUser_Success(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	svc := NewUserService(
		&mockUserRepo{
			getUser: &model.User{
				PasswordHash: string(hash),
				Email:        "test@test.com",
			},
			getErr: nil, createErr: nil},
		&mockTokenManager{},
		&mockSessionRepo{session: &model.Session{}},
	)

	accessToken, refreshToken, err := svc.LoginUser(context.Background(), "test@test.com", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "access_token", accessToken)
	assert.Equal(t, "refresh_token", refreshToken)
}
