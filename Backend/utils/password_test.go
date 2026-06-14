package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_RetornaHashDistinto(t *testing.T) {
	password := "123456"

	hash, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestCheckPassword_PasswordCorrecta_RetornaTrue(t *testing.T) {
	password := "123456"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	resultado := CheckPassword(password, hash)

	assert.True(t, resultado)
}

func TestCheckPassword_PasswordIncorrecta_RetornaFalse(t *testing.T) {
	password := "123456"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	resultado := CheckPassword("incorrecta", hash)

	assert.False(t, resultado)
}
