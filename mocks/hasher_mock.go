package mocks

import "webapi/pkg/app/interfaces"

type hasherMocked struct{}

func (m hasherMocked) Hahser(text string) (string, error) {
	return "", nil
}

func (m hasherMocked) Verify(originalText, hashedText string) bool {
	return true
}

func NewHasherMock() interfaces.IHasher {
	return &hasherMocked{}
}
