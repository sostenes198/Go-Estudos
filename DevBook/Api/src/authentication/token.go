package authentication

import (
	"devbook/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func CreateToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.TokenSecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeyValidation)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyValidation(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.TokenSecretKey, nil
}
