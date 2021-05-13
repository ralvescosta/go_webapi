package presentation

import (
	"github.com/gin-gonic/gin"

	httphandlers "webapi/pkg/interface/http_handlers"
)

func NewUserRoutes(router *gin.Engine, handler httphandlers.IUserHTTPHandler) {
	router.POST("/users", handler.Create)
	router.GET("/users", handler.GetAll)
	router.GET("/users/{id}", handler.GetById)
	router.PUT("/users/{id}", handler.UpdateById)
	router.DELETE("/users/{id}", handler.DeleteById)
}
