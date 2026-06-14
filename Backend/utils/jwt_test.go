package utils

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerarJWT_RetornaTokenValido(t *testing.T) {
	tokenString, err := GenerarJWT(1, "cliente")

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	assert.NoError(t, err)
	assert.True(t, token.Valid)
}
