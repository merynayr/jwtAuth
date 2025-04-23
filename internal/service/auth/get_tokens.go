package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/utils/hash"
	"github.com/merynayr/jwtauth/internal/utils/jwt"
)

// IssueTokens создаёт пару access + refresh токенов и сохраняет refresh
func (s *srv) IssueTokens(ctx context.Context, req model.AuthRequest) (*model.AuthResponse, error) {
	jwtClaims := &model.UserClaims{
		UserID: req.UserID,
	}

	accessToken, accessJTI, err := jwt.GenerateAccessToken(jwtClaims, s.authCfg.AccessTokenSecretKey(), s.authCfg.AccessTokenExp())
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	hash, err := hash.Hash(refreshToken)
	if err != nil {
		return nil, err
	}

	refreshTokenObject := &model.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		TokenHash: hash,
		AccessJTI: accessJTI,
		IP:        req.IP,
		UserAgent: req.UserAgent,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(s.authCfg.RefreshTokenExp()),
	}

	_, err = s.authRepository.CreateRefreshToken(ctx, refreshTokenObject)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
