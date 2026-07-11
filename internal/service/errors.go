package service

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidUserID      = errors.New("invalid user ID")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidTaskID      = errors.New("invalid task ID")
	ErrInvalidTask        = errors.New("invalid task")
	ErrSessionExpired     = errors.New("session expired")
)
