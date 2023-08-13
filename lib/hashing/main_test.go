package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashing(t *testing.T) {
	t.Run("it hashes a string, and compares the hash with original string", func(t *testing.T) {
		testString := "test"
		hash, err := Hash(testString)
		assert.NotEmpty(t, hash)
		assert.Nil(t, err)

		result := CheckCode(testString, hash)
		assert.NotEmpty(t, result)
		assert.Nil(t, err)
	})
}
