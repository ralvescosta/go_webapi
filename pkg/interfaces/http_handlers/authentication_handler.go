package httphandlers

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/services"
	interfaces "webapi/pkg/interfaces"

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
		interfaces.BadRequest(c, err.Error())
		return
	}

	if err := authDto.Validate(); err != nil {
		interfaces.BadRequest(c, err.Error())
		return
	}

	authenticatedUser, err := h.authenticationUserService.Perform(authDto.Email, authDto.Password)
	if err != nil {
		switch err.(type) {
		case *errors.InternalError:
			interfaces.InternalServerError(c, err.Error())
			return
		case *errors.NotFoundError:
			interfaces.NotFound(c, "User Not Found")
			return
		case *errors.UnauthorizeError:
			interfaces.Unauthorized(c, err.Error())
			return
		default:
			interfaces.InternalServerError(c, "Some internal Error occur, try again latter!")
			return
		}
	}

	interfaces.Ok(c, authenticatedUser)
}

func NewAuthenticationHTTPHandler(service services.IAuthenticationUser) IAuthenticationHTTPHandler {
	return &authenticationHTTPHandler{
		authenticationUserService: service,
	}
}
