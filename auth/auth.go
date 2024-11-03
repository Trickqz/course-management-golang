package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("sua_chave_secreta")

func GerarToken(usuarioID int) (string, error) {
	claims := jwt.MapClaims{
		"usuario_id": strconv.Itoa(usuarioID),
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidarToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
