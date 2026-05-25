package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("clave_secreta_temporal")

func GenerarJWT(usuarioID uint, rol string) (string, error) {
	claims := jwt.MapClaims{
		"usuario_id": usuarioID,
		"rol":        rol,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
