package hasher

import (
	"testing"
	"webapi/pkg/app/interfaces"

	"github.com/stretchr/testify/assert"
)

var hahser interfaces.IHasher = NewHahser()

func TestHahser_Hahser(t *testing.T) {
	_, err := hahser.Hahser("text")

	assert.Equal(t, err, nil)
}

func TestHahser_Verify(t *testing.T) {
	hash, err := hahser.Hahser("text")
	result := hahser.Verify("text", hash)

	assert.Equal(t, err, nil)
	assert.True(t, result)
}
