package httphandlers

import (
	"net/http/httptest"
	"testing"
	"webapi/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_ShouldReturnIUserHTTPHandlerInstance(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	userHandler := NewUserHTTPHandler(userServiceMock)

	assert.Implements(t, new(IUserHTTPHandler), userHandler)
}

func TestUserHandler_Create(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.Create(contextMock)
}
