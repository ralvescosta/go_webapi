package mocks

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/interfaces"
)

type tokenManagerMocked struct{}

func (m tokenManagerMocked) GenerateToken(tokenData *dtos.TokenDataDto) (string, error) {
	return "", nil
}
func (m tokenManagerMocked) VerifyToken(token string) (interface{}, error) {
	return "", nil
}

func NewTokenManagerMock() interfaces.ITokenManager {
	return &tokenManagerMocked{}
}
