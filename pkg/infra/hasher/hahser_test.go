package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"webapi/mocks"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
)

var hahser interfaces.IHasher = NewHahser()

func TestHahser_Hahser(t *testing.T) {
	var bcryptMock = mocks.NewBcryptMock(false, []byte("hashed"))
	generateHash = bcryptMock.GenerateHash

	hashed, err := hahser.Hahser("text")

	assert.Equal(t, err, nil)
	assert.Equal(t, hashed, "hashed")
}

func TestHahser_ShouldReturnErrorIfSameErrorOccurInCrypto(t *testing.T) {
	var bcryptMock = mocks.NewBcryptMock(true, []byte("hash"))
	generateHash = bcryptMock.GenerateHash

	_, err := hahser.Hahser("text")

	assert.Error(t, err)
	assert.IsType(t, err, errors.InternalError{})
}

func TestHahser_Verify(t *testing.T) {
	var bcryptMock = mocks.NewBcryptMock(false, []byte(""))
	compareHash = bcryptMock.CompareHash

	result := hahser.Verify("text", "hash")

	assert.True(t, result)
}

func TestHahser_ShouldRetorneFalseIfHashIsWrong(t *testing.T) {
	var bcryptMock = mocks.NewBcryptMock(true, []byte(""))
	compareHash = bcryptMock.CompareHash

	result := hahser.Verify("text", "hash")

	assert.False(t, result)
}
