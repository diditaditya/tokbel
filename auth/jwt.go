package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id int
	jwt.RegisteredClaims
}

func CreateJWT(id int) (string, error) {
	rawKey := os.Getenv("JWT_SECRET")
	signingKey := []byte(rawKey)

	expirationTime := time.Now().Add(24 * 60 * time.Minute)
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyJWT(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("unknown error")
	}

	return claims.Id, nil
}
