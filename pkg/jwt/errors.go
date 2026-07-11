package jwt

import "errors"

var (
	ErrTokenExpired   = errors.New("token expired")
	ErrTokenInvalid   = errors.New("invalid token")
	ErrEmptySignature = errors.New("empty signature key")
	ErrInvalidTTL     = errors.New("invalid TTL")
)
