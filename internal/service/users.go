package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/aidostt/task-manager/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo           repository.UserRepository
	sessionManager jwt.TokenManager
	sessionRepo    repository.SessionRepository
}

func NewUserService(repo repository.UserRepository, sManager jwt.TokenManager, sRepo repository.SessionRepository) User {
	return &userService{repo: repo, sessionManager: sManager, sessionRepo: sRepo}
}

func (s *userService) RegisterUser(ctx context.Context, email, plainPassword string) (string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	if user != nil {
		return "", "", ErrUserAlreadyExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	user = &model.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	user.ID, err = s.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return "", "", ErrUserAlreadyExists
		}
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	return s.createSession(ctx, user.ID)
}

func (s *userService) LoginUser(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", "", ErrInvalidCredentials
		}
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	if user == nil {
		return "", "", ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", ErrInvalidCredentials
		}
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	return s.createSession(ctx, user.ID)
}

func (s *userService) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	session, err := s.sessionRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	if session.ExpiresAt.Before(time.Now()) {
		return "", "", ErrSessionExpired
	}
	err = s.sessionRepo.DeleteSession(ctx, refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	return s.createSession(ctx, session.UserID)
}

func (s *userService) createSession(ctx context.Context, userID uuid.UUID) (string, string, error) {
	accessToken, err := s.sessionManager.NewAccessToken(userID.String())
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	refreshToken, err := s.sessionManager.NewRefreshToken()
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	session := &model.Session{
		UserID:    userID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.sessionManager.RefreshTTL()),
	}
	err = s.sessionRepo.CreateSession(ctx, session)
	if err != nil {
		return "", "", fmt.Errorf("internal error: %w", err)
	}
	return accessToken, refreshToken, nil
}
