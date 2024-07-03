package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("secretkey")

type ClaimKey struct {
	ID uint
	jwt.StandardClaims
}

func GenerateJWT(id uint) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &ClaimKey{
		id,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtkey)
	return
}

func ValidateToken(singnedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		singnedToken,
		&ClaimKey{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*ClaimKey)
	if !ok {
		err = errors.New("couldnt pass claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expires")
		return
	}
	return
}
