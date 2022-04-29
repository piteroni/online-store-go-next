package auth

import (
	common "piteroni/online-store-api"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGenerator struct{}

const expire = time.Hour * 168 // 1 week

func (JWTTokenGenerator) GenerateJWTToken(userID uint) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(expire).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret, err := common.Env("JWT_SECRET_KEY")
	if err != nil {
		return "", err
	}

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
