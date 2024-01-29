package pkg

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.MapClaims) (token string, err error) {

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtClaims.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return
	}

	return
}

func ValidateJWTAndReturnClaims(tokenStr string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("token not valid")
		return
	}
	return
}
