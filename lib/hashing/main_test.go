package hashing_test

import (
	"testing"

	"github.com/Goldziher/go-monorepo/lib/hashing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	t.Run("it hashes a string, and compares the hash with original string", func(t *testing.T) {
		testString := "test"
		hash, err := hashing.Hash(testString)
		assert.NotEmpty(t, hash)
		assert.Nil(t, err)

		result := hashing.CheckCode(testString, hash)
		assert.NotEmpty(t, result)
		assert.Nil(t, err)
	})
}
