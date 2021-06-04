package mocks

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JwtMocked struct {
	fileReaderFailure bool
	parseRSAFailure   bool
	claimsGenFailure  bool
}

func (m JwtMocked) FileReader(path string) ([]byte, error) {
	if m.fileReaderFailure {
		return nil, errors.New("file reader error")
	}

	return []byte(""), nil
}

func (m JwtMocked) ParseRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	if m.parseRSAFailure {
		return nil, errors.New("parse rsa private key")
	}

	return rsa.GenerateKey(rand.Reader, 2048)
}

func (m JwtMocked) ClaimsGenerator(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
	if m.claimsGenFailure {
		return nil
	}

	return &jwt.Token{
		Header: map[string]interface{}{},
		Claims: jwt.MapClaims{},
		Method: method,
	}
}

func NewJwtMock(fileReaderFailure, parseRSAFailure, claimsGenFailure bool) *JwtMocked {
	return &JwtMocked{
		fileReaderFailure: fileReaderFailure,
		parseRSAFailure:   parseRSAFailure,
		claimsGenFailure:  claimsGenFailure,
	}
}
