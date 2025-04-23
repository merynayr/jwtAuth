package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/jwtauth/internal/sys"
)

// RefreshTokens  выполняет Refresh операцию на пару Access, Refresh токенов
// @Summary      Обновить пару токенов
// @Description  Обновляет Access и Refresh токены. Refresh токен может быть передан через Cookie или Header "refresh_token"
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} model.AuthResponse
// @Failure      400 {object} sys.ErrorResponse
// @Failure      401 {object} sys.ErrorResponse
// @Router       /auth/refresh [post]
// @Param        refresh_token header string false "Refresh Token (если не используется Cookie)"
// @Param        access_token header string false "Access Token (если не используется Cookie)"
func (a *API) RefreshTokens(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		refreshToken = c.GetHeader("refresh_token")
		if refreshToken == "" {
			sys.HandleError(c, err)
			return
		}
	}

	accessToken, err := c.Cookie("access_token")
	if err != nil {
		accessToken = c.GetHeader("access_token")
		if accessToken == "" {
			sys.HandleError(c, err)
			return
		}
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	tokenPair, err := a.authService.RefreshTokens(c.Request.Context(), refreshToken, accessToken, ip, userAgent)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	a.setCookies(c, tokenPair.RefreshToken, tokenPair.AccessToken)
	c.JSON(http.StatusOK, tokenPair)
}
