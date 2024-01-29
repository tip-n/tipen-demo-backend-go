package middleware

import (
	"errors"
	"strings"
	"tipen-demo/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthorizeUserAndReturnID(c *fiber.Ctx) (
	id int,
	err error,
) {
	headers := c.GetReqHeaders()
	authorizationHeader := headers["Authorization"]
	if len(authorizationHeader) == 0 {
		err = errors.New("bearer token not found")
		return
	}
	bearerToken := authorizationHeader[0]
	splitted := strings.Split(bearerToken, " ")
	if len(splitted) == 0 {
		err = errors.New("bearer token not valid")
		return
	}
	token := splitted[1]

	claims, err := pkg.ValidateJWTAndReturnClaims(token)
	if err != nil {
		return
	}

	if claims["user_id"] == nil {
		err = errors.New("token not valid")
		return
	}

	id = int(claims["user_id"].(float64))
	return
}

func (m *Middleware) AuthorizeSellerAndReturnID(c *fiber.Ctx) (
	id int,
	err error,
) {
	headers := c.GetReqHeaders()
	authorizationHeader := headers["Authorization"]
	if len(authorizationHeader) == 0 {
		err = errors.New("bearer token not found")
		return
	}
	bearerToken := authorizationHeader[0]
	splitted := strings.Split(bearerToken, " ")
	if len(splitted) == 0 {
		err = errors.New("bearer token not valid")
		return
	}
	token := splitted[1]

	claims, err := pkg.ValidateJWTAndReturnClaims(token)
	if err != nil {
		return
	}

	if claims["seller_id"] == nil {
		err = errors.New("token not valid")
		return
	}

	id = int(claims["seller_id"].(float64))
	return
}
