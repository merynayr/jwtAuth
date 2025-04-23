package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthRequest структура данных для создания токенов
type AuthRequest struct {
	UserID    string
	IP        string
	UserAgent string
}

// AuthResponse структура ответа с токенами
type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// RefreshToken структура refresh токена
type RefreshToken struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	TokenHash string    `db:"token_hash"`
	AccessJTI string    `db:"access_jti"`
	IP        string    `db:"ip"`
	UserAgent string    `db:"user_agent"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

// UserClaims структура claims jwt-токена
type UserClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
