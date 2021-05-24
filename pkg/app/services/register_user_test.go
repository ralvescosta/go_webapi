package services

import (
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"

	"github.com/stretchr/testify/assert"
)

var mockedUserDto dtos.UserDto = dtos.UserDto{
	FirstName: "Name",
	LastName:  "Name",
	Email:     "email@email.com",
	Password:  "password",
}

var mockedUserEntity entities.User = entities.User{}

func TestRegister_CreateUserSuccessfully(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(&mockedUserDto)

	assert.Nil(t, result)
}

func TestRegister_CreateUserWhenHasherFailure(t *testing.T) {
	hahserMock := mocks.NewHasherMock(true, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(&mockedUserDto)

	assert.Error(t, result)
}

func TestRegister_CreateUserWhenUserRepoFailure(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(true, &mockedUserEntity)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(&mockedUserDto)

	assert.Error(t, result)
}
