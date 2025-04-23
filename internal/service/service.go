package service

import (
	"context"

	"github.com/merynayr/jwtauth/internal/model"
)

// UserService интерфейс сервисного слоя user
type UserService interface {
	CreateUser(ctx context.Context) (*model.User, error)
}

// AuthService интерфейс сервисного слоя auth
type AuthService interface {
	IssueTokens(ctx context.Context, req model.AuthRequest) (*model.AuthResponse, error)
	RefreshTokens(
		ctx context.Context,
		refreshTokenRaw string,
		accessTokenRaw string,
		ip string,
		userAgent string,
	) (*model.AuthResponse, error)
}
