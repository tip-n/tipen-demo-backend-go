package route

import (
	"strings"
	"tipen-demo/handler"

	"github.com/gofiber/fiber/v2"
)

type RegisterUserBodyRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

type LoginUserBodyRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserBodyRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (r *Route) InitUser() {
	api := r.Router.Group("/v1")
	user := api.Group("/user")
	user.Post("/register", func(c *fiber.Ctx) error {
		p := RegisterUserBodyRequest{}

		if err := c.BodyParser(&p); err != nil {
			return err
		}

		// Validation
		if errs := r.Validator.Validate(p); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, r.Validator.ErrorMessage(err))
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, ", "),
			}
		}

		ID, err := r.Handler.RegisterUser(handler.RegisterUserParams{
			Firstname: p.Firstname,
			Lastname:  p.Lastname,
			Password:  p.Password,
			Email:     p.Email,
		})
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrInternalServerError.Code,
				Message: err.Error(),
			}
		}
		return c.JSON(fiber.Map{
			"id": ID,
		})
	})

	user.Post("/login", func(c *fiber.Ctx) error {
		p := LoginUserBodyRequest{}

		if err := c.BodyParser(&p); err != nil {
			return err
		}

		if errs := r.Validator.Validate(p); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, r.Validator.ErrorMessage(err))
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, ", "),
			}
		}

		accessToken, err := r.Handler.LoginUser(handler.LoginUserParams{
			Email:    p.Email,
			Password: p.Password,
		})
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrInternalServerError.Code,
				Message: err.Error(),
			}
		}
		return c.JSON(fiber.Map{
			"access_token": accessToken,
		})
	})
	user.Put("/update", func(c *fiber.Ctx) error {
		p := UpdateUserBodyRequest{}

		userID, err := r.Middleware.AuthorizeUserAndReturnID(c)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: err.Error(),
			}
		}

		if err := c.BodyParser(&p); err != nil {
			return err
		}

		if errs := r.Validator.Validate(p); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, r.Validator.ErrorMessage(err))
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, ", "),
			}
		}
		if err := r.Handler.UpdateUser(handler.UpdateUserParams{
			ID:        userID,
			Firstname: p.Firstname,
			Lastname:  p.Lastname,
		}); err != nil {
			return err
		}
		return c.SendString("Successfully Updated")
	})
	user.Get("/profile", func(c *fiber.Ctx) error {
		userID, err := r.Middleware.AuthorizeUserAndReturnID(c)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: err.Error(),
			}
		}
		resp, err := r.Handler.GetUserProfile(userID)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrInternalServerError.Code,
				Message: err.Error(),
			}
		}
		return c.JSON(resp)
	})
}
