package services

import (
	"context"
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/errors"

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

	result := registerUserService.Register(context.Background(), &mockedUserDto)

	assert.Nil(t, result)
}

func TestRegister_ShouldReturnAlreadyExistErrorIfUserAlreadyExist(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	user := &entities.User{
		Email: "email@email.com",
	}
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, user)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(context.Background(), &mockedUserDto)

	assert.Error(t, result)
	assert.IsType(t, result, errors.AlreadyExisteError{})
}

func TestRegister_CreateUserWhenHasherFailure(t *testing.T) {
	hahserMock := mocks.NewHasherMock(true, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(context.Background(), &mockedUserDto)

	assert.Error(t, result)
}

func TestRegister_CreateUserWhenUserRepoFailure(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(true, &mockedUserEntity)
	registerUserService := NewUserService(userRepositoryInMemory, hahserMock)

	result := registerUserService.Register(context.Background(), &mockedUserDto)

	assert.Error(t, result)
}
