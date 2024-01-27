package pkg

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int64) (token string, err error) {
	t := time.Now()
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"iat":     t.Unix(),
		"exp":     t.Add(time.Hour * 24).Unix(),
	})

	token, err = jwtClaims.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return
	}

	return
}

func ValidateJWTAndGetID(tokenStr string) (id int, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["user_id"] == nil {
			err = errors.New("token not valid")
			return
		}
		id = int(claims["user_id"].(float64))
		return
	} else {
		err = errors.New("token not valid")
		return
	}
}
