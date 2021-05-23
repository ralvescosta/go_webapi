package mocks

import (
	"errors"
	"webapi/pkg/app/interfaces"
)

type hasherMocked struct {
	failure bool
}

func (m hasherMocked) Hahser(text string) (string, error) {
	if m.failure {
		return "", errors.New("Error")
	}

	return "", nil
}

func (m hasherMocked) Verify(originalText, hashedText string) bool {
	return true
}

func NewHasherMock(failure bool) interfaces.IHasher {
	return &hasherMocked{failure: failure}
}
