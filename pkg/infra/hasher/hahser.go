package hasher

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
)

type hasher struct{}

var generateHash = bcrypt.GenerateFromPassword
var compareHash = bcrypt.CompareHashAndPassword

func (h hasher) Hahser(text string) (string, error) {
	hashed, err := generateHash([]byte(text), 9)
	if err != nil {
		log.Errorf("hasher.Hahser - generate password hash: %v", err)
		return "", errors.NewInternalError(err.Error())
	}

	return string(hashed), nil
}

func (h hasher) Verify(originalText, hashedText string) bool {
	if err := compareHash([]byte(hashedText), []byte(originalText)); err != nil {
		return false
	}

	return true
}

func NewHahser() interfaces.IHasher {
	return &hasher{}
}
