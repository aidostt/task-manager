package jwt

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewAccessToken(string) (string, error)
	NewRefreshToken() (string, error)
	Parse(string) (*Claims, error)
	RefreshTTL() time.Duration
}

type Manager struct {
	signingKey string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

type Claims struct {
	UserID    string   `json:"user_id"`
	Roles     []string `json:"roles,omitempty"`
	Activated bool     `json:"activated,omitempty"`
	jwt.RegisteredClaims
}

func NewManager(signingKey string, accessTTL, refreshTTL time.Duration) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	if accessTTL < 0 || refreshTTL < 0 {
		return nil, errors.New("invalid access TTL or refresh TTL")
	}
	return &Manager{signingKey: signingKey, accessTTL: accessTTL, refreshTTL: refreshTTL}, nil
}

func (m *Manager) RefreshTTL() time.Duration {
	return m.refreshTTL
}

func (m *Manager) NewAccessToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

// Parse taking from the payload of JWT user id and returns it in string format. Token is still returned
// in both cases, if it is expired or not.
func (m *Manager) Parse(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
