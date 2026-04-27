package auth

import (
	"edge-gateway/internal/apperr"
	"fmt"
	"time"
)

type Claims struct {
	UserID    string
	ExpiresAt time.Time
}

var staticTokens = map[string]Claims{
	"token-user-1": {
		UserID:    "USER-1",
		ExpiresAt: time.Date(2099, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	"token-user-2": {
		UserID:    "USER-2",
		ExpiresAt: time.Date(2099, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	"expired-user-1": {
		UserID:    "USER-1",
		ExpiresAt: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

func ValidateToken(token string) (Claims, error) {
	if token == "" {
		return Claims{}, fmt.Errorf("%w: missing token", apperr.ErrUnauthorized)
	}

	claims, ok := staticTokens[token]
	if !ok {
		return Claims{}, fmt.Errorf("%w: invalid token %q", apperr.ErrUnauthorized, token)
	}

	if time.Now().After(claims.ExpiresAt) {
		return Claims{}, fmt.Errorf("%w: expired token %q", apperr.ErrUnauthorized, token)
	}

	return claims, nil
}
