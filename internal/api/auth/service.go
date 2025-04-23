package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/jwtauth/internal/config"
	"github.com/merynayr/jwtauth/internal/service"
)

// API auth структура
type API struct {
	authService service.AuthService
	authConfig  config.AuthConfig
}

// NewAPI возвращает новый объект имплементации API-слоя auth
func NewAPI(authService service.AuthService, authConfig config.AuthConfig) *API {
	return &API{
		authService: authService,
		authConfig:  authConfig,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("auth")
	{
		authGroup.GET("/tokens", api.GetTokens)
		authGroup.POST("/refresh", api.RefreshTokens)
	}
}

// setCookies устанавливают токены в куки
func (api *API) setCookies(c *gin.Context, refreshToken string, accessToken string) {
	if len(refreshToken) > 0 {
		c.SetCookie("refresh_token", refreshToken, int(api.authConfig.RefreshTokenExp()*2), "/", "", false, true)
	}
	if len(accessToken) > 0 {
		c.SetCookie("access_token", accessToken, int(api.authConfig.AccessTokenExp()*2), "/", "", false, true)
	}
}
