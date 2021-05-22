package services

import (
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"

	"github.com/stretchr/testify/assert"
)

var mockedUser dtos.UserDto = dtos.UserDto{
	FirstName: "Name",
	LastName:  "Name",
	Email:     "email@email.com",
	Password:  "password",
}

func TestSomething(t *testing.T) {
	hahserMock := mocks.NewHasherMock()
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory()
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(&mockedUser)

	assert.Nil(t, result)

}
