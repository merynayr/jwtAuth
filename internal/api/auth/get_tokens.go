package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/sys"
)

// GetTokens обрабатывает HTTP-запрос на получение access токена
// @Summary      Получить пару токенов
// @Description  Возвращает Access и Refresh токены по user_id
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user_id query string true "UUID пользователя"
// @Success      200 {object} model.AuthResponse
// @Failure      400 {object} sys.ErrorResponse
// @Failure      500 {object} sys.ErrorResponse
// @Router       /auth/tokens [get]
func (a *API) GetTokens(c *gin.Context) {
	userID := c.Query("user_id")
	_, err := uuid.Parse(userID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	req := &model.AuthRequest{
		UserID:    userID,
		IP:        ip,
		UserAgent: userAgent,
	}

	token, err := a.authService.IssueTokens(c.Request.Context(), *req)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	a.setCookies(c, token.RefreshToken, token.AccessToken)
	c.JSON(http.StatusOK, token)
}
