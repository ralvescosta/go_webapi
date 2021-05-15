package hasher

import (
	"golang.org/x/crypto/bcrypt"

	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
)

type hasher struct{}

func (h hasher) Hahser(text string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), 9)
	if err != nil {
		return "", errors.NewInternalError(err.Error())
	}

	return string(hashed), nil
}

func NewHahser() interfaces.IHasher {
	return &hasher{}
}
