package mocks

import (
	"errors"
	"webapi/pkg/app/interfaces"
)

type hasherMocked struct {
	failure      bool
	verifyResult bool
}

func (m hasherMocked) Hahser(text string) (string, error) {
	if m.failure {
		return "", errors.New("Error")
	}

	return "", nil
}

func (m hasherMocked) Verify(originalText, hashedText string) bool {
	return m.verifyResult
}

func NewHasherMock(failure, verifyResult bool) interfaces.IHasher {
	return &hasherMocked{failure: failure, verifyResult: verifyResult}
}
