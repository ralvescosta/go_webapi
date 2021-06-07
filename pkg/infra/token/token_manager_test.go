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
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey(false)
	claimsGenerator = jwtMock.ClaimsGenerator(false)

	token, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInReadRSAPrivateKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(true)

	_, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error when try to read rsa private key")
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInParseRSAPrivateKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey(true)

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

func TestCreateToken_ShouldVerifyATokenCorrectly(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(false, time.Now().Add(time.Hour).Unix())

	token, err := manager.VerifyToken("token")

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInReadRSAPublicKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(true)

	token, err := manager.VerifyToken("token")

	assert.Nil(t, token)
	assert.Equal(t, err.Error(), "error when try to read rsa public key")
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInParseRSAPublicKey(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(true)

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error when try to create rsa public key")
}

func TestCreateToken_ShouldReturnErrIfSomeErrorOccurWhenParseClaims(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(true, time.Now().Add(time.Hour).Unix())

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid token")
}

func TestCreateToken_ShouldReturnErrIfTokenExpired(t *testing.T) {
	jwtMock := mocks.NewJwtMock()
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(false, time.Now().Add(-time.Hour).Unix())

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "jwt is expired")
}
