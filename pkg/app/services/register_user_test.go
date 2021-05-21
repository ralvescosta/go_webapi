package services

import (
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	hahserMock := mocks.NewHasherMock()
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory()
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(&dtos.UserDto{})

	assert.Nil(t, result)

}
