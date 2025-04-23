package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/sys"
	"github.com/merynayr/jwtauth/internal/utils/hash"
	"github.com/merynayr/jwtauth/internal/utils/jwt"
)

func (s *srv) RefreshTokens(
	ctx context.Context,
	refreshTokenRaw string,
	accessTokenRaw string,
	ip string,
	userAgent string,
) (*model.AuthResponse, error) {
	accessClaims, err := jwt.VerifyToken(accessTokenRaw, s.authCfg.AccessTokenSecretKey())
	if err != nil {
		return nil, err
	}

	userID := accessClaims.UserID
	accessJTI := accessClaims.Id

	tokenEntry, err := s.authRepository.FindRefreshToken(ctx, accessJTI)
	if err != nil {
		return nil, sys.InvalidRefreshTokenError
	}

	err = hash.CompareHashAndPass(tokenEntry.TokenHash, refreshTokenRaw)
	if err != nil {
		return nil, sys.InvalidRefreshTokenError
	}

	if time.Now().After(tokenEntry.ExpiresAt) {
		return nil, sys.RefreshTokenExpiredError
	}

	if tokenEntry.IP != ip || tokenEntry.UserAgent != userAgent {
		// можно асинхронно отправить email пользователю, например с помощью брокера сообщений
		fmt.Println("Отправялю сообщение на email")
	}

	err = s.authRepository.DeleteByID(ctx, tokenEntry.AccessJTI)
	if err != nil {
		return nil, err
	}

	jwtClaims := &model.UserClaims{
		UserID: tokenEntry.UserID,
	}

	newAccessToken, newAccessJTI, err := jwt.GenerateAccessToken(jwtClaims, s.authCfg.AccessTokenSecretKey(), s.authCfg.AccessTokenExp())
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}
	hash, err := hash.Hash(newRefreshToken)
	if err != nil {
		return nil, err
	}

	refreshModel := &model.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userID,
		TokenHash: hash,
		AccessJTI: newAccessJTI,
		IP:        ip,
		UserAgent: userAgent,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(s.authCfg.RefreshTokenExp()),
	}
	_, err = s.authRepository.CreateRefreshToken(ctx, refreshModel)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}, nil
}
