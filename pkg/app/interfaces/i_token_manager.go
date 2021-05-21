package interfaces

import (
	"webapi/pkg/app/dtos"
)

type ITokenManager interface {
	GenerateToken(tokenData *dtos.TokenDataDto) (string, error)
	VerifyToken(token string) (interface{}, error)
}
