package mocks

import "errors"

type BcryptMocked struct {
	failure bool
	hash    []byte
}

func (m BcryptMocked) GenerateHash(password []byte, cost int) ([]byte, error) {
	if m.failure {
		return []byte(""), errors.New("bcrypt error")
	}

	return m.hash, nil
}

func (m BcryptMocked) CompareHash(hashedPassword, password []byte) error {
	if m.failure {
		return errors.New("bcrypt error")
	}

	return nil
}

func NewBcryptMock(failure bool, hash []byte) *BcryptMocked {
	return &BcryptMocked{
		failure: failure,
		hash:    hash,
	}
}
