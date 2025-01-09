package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ExtractUserIdFromJwtClaims(claims jwt.Claims) (uuid.UUID, error) {
	sub, err := claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.Parse(sub)
}
