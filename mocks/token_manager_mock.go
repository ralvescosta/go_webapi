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
func (m tokenManagerMocked) VerifyToken(token string) (interface{}, error) {
	if m.failure {
		return "", errors.New("Error")
	}
	return "", nil
}

func NewTokenManagerMock(failure bool) interfaces.ITokenManager {
	return &tokenManagerMocked{failure: failure}
}
