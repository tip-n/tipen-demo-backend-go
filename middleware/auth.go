package middleware

import (
	"errors"
	"strings"
	"tipen-demo/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthorizeAndReturnID(c *fiber.Ctx) (
	ID int,
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

	ID, err = pkg.ValidateJWTAndGetID(token)
	return
}
