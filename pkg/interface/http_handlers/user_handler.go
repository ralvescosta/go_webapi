package httphandlers

import (
	"github.com/gin-gonic/gin"

	"webapi/pkg/app"
)

type IUserHTTPHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	UpdateById(c *gin.Context)
	DeleteById(c *gin.Context)
}

type userHTTPHandler struct {
	userService app.IUserService
}

func (h *userHTTPHandler) Create(c *gin.Context) {

}

func (h *userHTTPHandler) GetAll(c *gin.Context) {

}

func (h *userHTTPHandler) GetById(c *gin.Context) {

}

func (h *userHTTPHandler) UpdateById(c *gin.Context) {

}

func (h *userHTTPHandler) DeleteById(c *gin.Context) {

}

func NewUserHTTPHandler(userService app.IUserService) IUserHTTPHandler {
	return &userHTTPHandler{
		userService: userService,
	}
}
