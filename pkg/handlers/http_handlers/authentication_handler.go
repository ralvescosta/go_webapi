package httphandlers

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/services"
	handlers "webapi/pkg/handlers"
	"webapi/pkg/handlers/validators"

	"github.com/gin-gonic/gin"
)

type IAuthenticationHTTPHandler interface {
	Create(c *gin.Context)
}

type authenticationHTTPHandler struct {
	authenticationUserService services.IAuthenticationUser
}

func (h authenticationHTTPHandler) Create(c *gin.Context) {
	authDto := &dtos.AuthenticationDTO{}

	if err := c.ShouldBindJSON(authDto); err != nil {
		handlers.BadRequest(c, err.Error())
		return
	}

	messages, err := validators.ValidateAuthUserBody(authDto)
	if err != nil {
		handlers.InvalidBody(c, messages)
		return
	}

	authenticatedUser, err := h.authenticationUserService.Perform(authDto.Email, authDto.Password)
	if err != nil {
		switch err.(type) {
		case errors.InternalError:
			handlers.InternalServerError(c, err.Error())
			return
		case errors.NotFoundError:
			handlers.NotFound(c, "User Not Found")
			return
		case errors.UnauthorizeError:
			handlers.Unauthorized(c, err.Error())
			return
		case errors.BadRequestError:
			handlers.Unauthorized(c, err.Error())
			return
		default:
			handlers.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		}
	}

	handlers.Ok(c, authenticatedUser)
}

func NewAuthenticationHTTPHandler(service services.IAuthenticationUser) IAuthenticationHTTPHandler {
	return &authenticationHTTPHandler{
		authenticationUserService: service,
	}
}
