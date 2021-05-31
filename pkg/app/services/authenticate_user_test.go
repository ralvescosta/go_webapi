package services

import (
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"

	"github.com/stretchr/testify/assert"
)

var mockedAuthenticationDto dtos.AuthenticationDTO = dtos.AuthenticationDTO{
	Email:    "email@email.com",
	Password: "password",
}

func TestPerform_ShouldAuthenticateUserSuccessfully(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	mockedUserEntity.Email = "email@email.com"
	mockedUserEntity.Password = "123456"
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock(false)
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, _ := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password, "")

	assert.NotNil(t, result)

	mockedUserEntity.Email = ""
	mockedUserEntity.Password = ""
}

func TestPerform_ShouldReturnInternalErrorIfOccurErrorInRepository(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(true, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock(false)
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, err := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password, "")

	assert.Nil(t, result)
	assert.IsType(t, err, errors.InternalError{})
}

func TestPerform_ShouldReturnBadRequestErrorIfUserDontExist(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock(false)
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, err := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password, "")

	assert.Nil(t, result)
	assert.IsType(t, err, errors.BadRequestError{})
}

func TestPerform_ShouldReturnBadRequestIfUserPasswordIsWrong(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, false)
	mockedUserEntity.Email = "email@email.com"
	mockedUserEntity.Password = "123456"
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock(false)
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, err := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password, "")

	assert.Nil(t, result)
	assert.IsType(t, err, errors.BadRequestError{})

	mockedUserEntity.Email = ""
	mockedUserEntity.Password = ""
}

func TestPerform_ShouldReturnInternalErrorIfSomeErrorOccurInTokenGeneration(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false, true)
	mockedUserEntity.Email = "email@email.com"
	mockedUserEntity.Password = "123456"
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock(true)
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, err := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password, "")

	assert.Nil(t, result)
	assert.IsType(t, err, errors.InternalError{})

	mockedUserEntity.Email = ""
	mockedUserEntity.Password = ""
}
