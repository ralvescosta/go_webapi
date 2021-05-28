package hasher

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
)

type hasher struct{}

func (h hasher) Hahser(text string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), 9)
	if err != nil {
		err = fmt.Errorf("hasher.Hahser - generate password hash: %v", err)
		log.Print(err)
		return "", errors.NewInternalError(err.Error())
	}

	return string(hashed), nil
}

func (h hasher) Verify(originalText, hashedText string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(originalText)); err != nil {
		return false
	}

	return true
}

func NewHahser() interfaces.IHasher {
	return &hasher{}
}
