package httphandlers

import (
	"net/http/httptest"
	"testing"
	"webapi/mocks"

	"github.com/gin-gonic/gin"
)

func TestCreate(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.Create(contextMock)
}

func TestGetById(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.GetById(contextMock)
}

func TestGetByAll(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.GetAll(contextMock)
}

func TestUpdateById(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.UpdateById(contextMock)
}

func TestDeleteById(t *testing.T) {
	userServiceMock := mocks.NewUserServicesMock()
	handler := userHTTPHandler{
		userService: userServiceMock,
	}

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())

	handler.DeleteById(contextMock)
}
