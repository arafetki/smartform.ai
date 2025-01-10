package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func HMACCheck(tokenString string, secret string) (string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing algorithm: %s", t.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	sub, err := parsedToken.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return sub, nil
}
