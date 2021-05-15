package httphandlers

import (
	"github.com/gin-gonic/gin"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/services"
	interfaces "webapi/pkg/interfaces"
)

type IUserHTTPHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	UpdateById(c *gin.Context)
	DeleteById(c *gin.Context)
}

type userHTTPHandler struct {
	userService services.IUserService
}

func (h *userHTTPHandler) Create(c *gin.Context) {
	userDto := &dtos.UserDto{}

	if err := c.ShouldBindJSON(userDto); err != nil {
		interfaces.BadRequest(c, err.Error())
		return
	}

	if err := userDto.Validate(); err != nil {
		interfaces.BadRequest(c, err.Error())
		return
	}

	if err := h.userService.Register(userDto); err != nil {
		switch err.(type) {
		case *errors.InternalError:
			interfaces.InternalServerError(c, err.Error())
			return
		case *errors.AlreadyExisteError:
			interfaces.ConflictError(c, err.Error())
			return
		default:
			interfaces.InternalServerError(c, err.Error())
			return
		}
	}
	interfaces.Created(c)
}

func (h *userHTTPHandler) GetAll(c *gin.Context) {

}

func (h *userHTTPHandler) GetById(c *gin.Context) {

}

func (h *userHTTPHandler) UpdateById(c *gin.Context) {

}

func (h *userHTTPHandler) DeleteById(c *gin.Context) {

}

func NewUserHTTPHandler(userService services.IUserService) IUserHTTPHandler {
	return &userHTTPHandler{
		userService: userService,
	}
}
