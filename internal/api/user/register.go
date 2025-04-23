package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/merynayr/jwtauth/internal/sys"
)

// CreateUser создание пользователя и получение GUID пользователя
// @Summary Создать пользователя
// @Description Создаёт пользователя и возвращает GUID пользователя
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 400 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /register [post]
func (a *API) CreateUser(c *gin.Context) {
	GUID, err := a.userService.CreateUser(c.Request.Context())
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, GUID)
}
