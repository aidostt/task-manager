package service

import (
	"context"
	"errors"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *User {
	return &User{repo: repo}
}

func (s *User) RegisterUser(ctx context.Context, email, plainPassword string) error {
	if email == "" || plainPassword == "" {
		return errors.New("email or password is empty")
	}

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if user != nil {
		return errors.New("user with this email already exists")
	}
	// пользователь не найден — продолжаем регистрацию
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user = &model.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	err = s.repo.CreateUser(ctx, user)
	return err
}

func (s *User) LoginUser(ctx context.Context, email, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.New("email or password is empty")
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return "", "", errors.New("wrong password")
		default:
			return "", "", err
		}
	}
	//TODO: generate jwt tokens and return them

	return "", "", nil
}
