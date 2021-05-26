package presentation

import (
	"github.com/gin-gonic/gin"

	httphandlers "webapi/pkg/interfaces/http_handlers"
)

func NewUserRoutes(router *gin.Engine, handler httphandlers.IUserHTTPHandler) {
	router.POST("/users", handler.Create)
}

func NewAuthenticationRoute(router *gin.Engine, handler httphandlers.IAuthenticationHTTPHandler) {
	router.POST("/authentication", handler.Create)
}
