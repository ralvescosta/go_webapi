package token

import (
	"testing"
	"time"
	"webapi/mocks"
	"webapi/pkg/app/dtos"

	"github.com/stretchr/testify/assert"
)

var manager = NewTokenManager()
var tokenData = &dtos.TokenDataDto{
	Id:       1,
	ExpireIn: time.Now().Add(time.Hour),
	Audience: "audience",
}

func TestCreateToken_ShouldCreateAToken(t *testing.T) {
	jwtMock := mocks.NewJwtMock(false, false, false)
	fileReader = jwtMock.FileReader
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey
	claimsGenerator = jwtMock.ClaimsGenerator

	token, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInReadRSAPrivateKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock(true, false, false)
	fileReader = jwtMock.FileReader

	_, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error when try to read rsa private key")
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInParseRSAPrivateKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock(false, true, false)
	fileReader = jwtMock.FileReader
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey

	_, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error when try to create rsa private key")
}

// func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInClaimGenerator(t *testing.T) {
// 	jwtMock := mocks.NewJwtMock(false, false, true)
// 	fileReader = jwtMock.FileReader
// 	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey
// 	claimsGenerator = jwtMock.ClaimsGenerator

// 	_, err := manager.GenerateToken(tokenData)

// 	assert.NotNil(t, err)
// }
