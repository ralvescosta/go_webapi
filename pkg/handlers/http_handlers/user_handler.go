package httphandlers

import (
	"github.com/gin-gonic/gin"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/services"
	handlers "webapi/pkg/handlers"
	"webapi/pkg/handlers/validators"
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
		handlers.BadRequest(c, err.Error())
		return
	}

	messages, err := validators.ValidateCreateUserBody(userDto)
	if err != nil {
		handlers.InvalidBody(c, messages)
		return
	}

	if err := h.userService.Register(c, userDto); err != nil {
		switch err.(type) {
		case errors.InternalError:
			handlers.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		case errors.AlreadyExisteError:
			handlers.Conflict(c, "User Already Exist!")
			return
		default:
			handlers.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		}
	}
	handlers.Created(c)
}

func NewUserHTTPHandler(userService services.IUserService) IUserHTTPHandler {
	return &userHTTPHandler{
		userService: userService,
	}
}
