package jwtauth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var TokenSecret []byte

//Generate token...
func GenToken(username, password string) (string, error) {
	c := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "cloudcost-demo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(TokenSecret)
}

//Parse token...
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
