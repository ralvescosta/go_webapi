package presentation

import (
	httphandlers "webapi/pkg/interface/http_handlers"

	"github.com/gin-gonic/gin"
)

func NewUserRoutes(router *gin.Engine, handler httphandlers.IUserHTTPHandler) {
	router.POST("/users", handler.Create)
	router.GET("/users", handler.GetAll)
	router.GET("/users/{id}", handler.GetById)
	router.PUT("/users/{id}", handler.UpdateById)
	router.DELETE("/users/{id}", handler.DeleteById)
}
