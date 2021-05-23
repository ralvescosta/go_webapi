package services

import (
	"testing"
	"webapi/mocks"
	"webapi/pkg/app/dtos"

	"github.com/stretchr/testify/assert"
)

var mockedAuthenticationDto dtos.AuthenticationDTO = dtos.AuthenticationDTO{
	Email:    "email@email.com",
	Password: "password",
}

func TestPerform_AuthenticateUserSuccessfully(t *testing.T) {
	hahserMock := mocks.NewHasherMock(false)
	userRepositoryInMemory := mocks.NewUserRepositoryInMemory(false, &mockedUserEntity)
	tokenManagerMock := mocks.NewTokenManagerMock()
	authenticateUserService := NewAuthenticationUser(userRepositoryInMemory, hahserMock, tokenManagerMock)

	result, _ := authenticateUserService.Perform(mockedAuthenticationDto.Email, mockedAuthenticationDto.Password)

	assert.NotNil(t, result)
}
