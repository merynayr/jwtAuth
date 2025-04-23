package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/merynayr/jwtauth/internal/config"
)

const (
	accessTokenSecretKeyEnvName = "ACCESS_TOKEN_SECRET_KEY" // #nosec G101
	accessTokenExpEnvName       = "ACCESS_TOKEN_EXP"        // #nosec G101
	refreshTokenExpEnvName      = "REFRESH_TOKEN_EXP"       // #nosec G101
)

type authConfig struct {
	refreshTokenExp time.Duration

	accessTokenSecretKey []byte
	accessTokenExp       time.Duration
}

// NewAuthConfig returns new auth service config
func NewAuthConfig() (config.AuthConfig, error) {
	refreshTokenExp, err := strconv.Atoi(os.Getenv(refreshTokenExpEnvName))
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token exp. time")
	}

	accessTokenSecretKey := os.Getenv(accessTokenSecretKeyEnvName)
	if len(accessTokenSecretKey) == 0 {
		return nil, fmt.Errorf("access token secret not found")
	}

	accessTokenExp, err := strconv.Atoi(os.Getenv(accessTokenExpEnvName))
	if err != nil {
		return nil, fmt.Errorf("invalid access token exp. time")
	}

	return &authConfig{
		refreshTokenExp:      time.Minute * time.Duration(refreshTokenExp),
		accessTokenSecretKey: []byte(accessTokenSecretKey),
		accessTokenExp:       time.Minute * time.Duration(accessTokenExp),
	}, nil
}

func (cfg *authConfig) RefreshTokenExp() time.Duration {
	return cfg.refreshTokenExp
}

func (cfg *authConfig) AccessTokenSecretKey() []byte {
	return cfg.accessTokenSecretKey
}

func (cfg *authConfig) AccessTokenExp() time.Duration {
	return cfg.accessTokenExp
}
