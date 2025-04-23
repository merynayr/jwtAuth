package user

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/jwtauth/internal/service"
)

// API auth структура
type API struct {
	userService service.UserService
}

// NewAPI возвращает новый объект имплементации API-слоя auth
func NewAPI(userService service.UserService) *API {
	return &API{
		userService: userService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	router.POST("/register", api.CreateUser)
}
