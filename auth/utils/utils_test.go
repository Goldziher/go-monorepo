package utils_test

import (
	"testing"

	"github.com/Goldziher/go-monorepo/auth/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateRandomString(t *testing.T) {
	result := utils.CreateStateString()
	assert.NotEmpty(t, result)
}
