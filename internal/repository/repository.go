package repository

import (
	"context"

	"github.com/merynayr/jwtauth/internal/model"
)

// UserRepository - интерфейс репо слоя user
type UserRepository interface {
	CreateUser(ctx context.Context, req *model.User) (*model.User, error)
}

// AuthRepository - интерфейс репо слоя jwt
type AuthRepository interface {
	CreateRefreshToken(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error)
	DeleteByID(ctx context.Context, jti string) error
	FindRefreshToken(ctx context.Context, jti string) (*model.RefreshToken, error)
}
