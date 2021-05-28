package httphandlers

import (
	"net/http/httptest"
	"testing"
	"webapi/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnIAuthenticationHTTPHandlere(t *testing.T) {
	authServiceMock := mocks.NewAuthenticationUserMocked()
	authHandler := NewAuthenticationHTTPHandler(authServiceMock)

	assert.Implements(t, new(IAuthenticationHTTPHandler), authHandler)
}

func TestAuthenticationHandler_Create(t *testing.T) {
	authServiceMock := mocks.NewAuthenticationUserMocked()
	handler := authenticationHTTPHandler{
		authenticationUserService: authServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.Create(contextMock)
}
