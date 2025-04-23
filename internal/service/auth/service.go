package auth

import (
	"github.com/merynayr/jwtauth/internal/config"
	"github.com/merynayr/jwtauth/internal/repository"
	"github.com/merynayr/jwtauth/internal/service"
)

type srv struct {
	authRepository repository.AuthRepository
	authCfg        config.AuthConfig
}

// NewService возвращает новый объект сервисного слоя jwtauth
func NewService(authRepo repository.AuthRepository, authCfg config.AuthConfig) service.AuthService {
	return &srv{
		authRepository: authRepo,
		authCfg:        authCfg,
	}
}
