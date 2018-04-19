package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetToken(signInKey string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0,0,1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signInKey))
}