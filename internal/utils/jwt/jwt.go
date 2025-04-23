package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/sys"

	"github.com/dgrijalva/jwt-go"
)

// GenerateAccessToken генерирует jwt-токен
func GenerateAccessToken(claims *model.UserClaims, secretKey []byte, duration time.Duration) (string, string, error) {
	if claims == nil {
		return "", "", fmt.Errorf("info is nil")
	}
	jti := uuid.NewString()
	claims.StandardClaims = jwt.StandardClaims{
		Subject:   claims.UserID,
		Id:        jti,
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signed, err := token.SignedString(secretKey)
	return signed, jti, err
}

// GenerateRefreshToken генерирует refresh токен
func GenerateRefreshToken() (raw string, err error) {
	buf := make([]byte, 54)
	_, err = rand.Read(buf)
	if err != nil {
		return "", err
	}
	raw = base64.StdEncoding.EncodeToString(buf)
	return raw, nil
}

// VerifyToken валидирует jwt-токен и возвращает его claims
func VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			return nil, sys.AccessTokenExpiredError
		}
		return nil, fmt.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
