package mocks

import (
	"errors"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/interfaces"
)

type tokenManagerMocked struct {
	failure bool
}

func (m tokenManagerMocked) GenerateToken(tokenData *dtos.TokenDataDto) (string, error) {
	if m.failure {
		return "", errors.New("Error")
	}
	return "", nil
}
func (m tokenManagerMocked) VerifyToken(token string) (*dtos.AuthenticatedUserDto, error) {
	if m.failure {
		return nil, errors.New("Error")
	}
	return nil, nil
}

func NewTokenManagerMock(failure bool) interfaces.ITokenManager {
	return &tokenManagerMocked{failure: failure}
}
