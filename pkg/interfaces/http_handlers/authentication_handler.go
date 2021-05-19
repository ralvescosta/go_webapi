package httphandlers

import (
	"webapi/pkg/app/services"

	"github.com/gin-gonic/gin"
)

type IAuthenticationHTTPHandler interface {
	Create(c *gin.Context)
}

type authenticationHTTPHandler struct {
	authenticationUserService services.IAuthenticationUser
}

func (h authenticationHTTPHandler) Create(c *gin.Context) {}

func NewAuthenticationHTTPHandler(service services.IAuthenticationUser) IAuthenticationHTTPHandler {
	return &authenticationHTTPHandler{
		authenticationUserService: service,
	}
}
