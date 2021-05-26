package httphandlers

import (
	"github.com/gin-gonic/gin"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/services"
	interfaces "webapi/pkg/interfaces"
	"webapi/pkg/interfaces/validators"
)

type IUserHTTPHandler interface {
	Create(c *gin.Context)
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

	messages, err := validators.ValidateCreateUserBody(userDto)
	if err != nil {
		interfaces.InvalidBody(c, messages)
		return
	}

	if err := h.userService.Register(userDto); err != nil {
		switch err.(type) {
		case errors.InternalError:
			interfaces.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		case errors.AlreadyExisteError:
			interfaces.Conflict(c, "User Already Exist!")
			return
		default:
			interfaces.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		}
	}
	interfaces.Created(c)
}

func NewUserHTTPHandler(userService services.IUserService) IUserHTTPHandler {
	return &userHTTPHandler{
		userService: userService,
	}
}
