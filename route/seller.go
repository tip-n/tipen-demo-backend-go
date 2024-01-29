package route

import (
	"strings"
	"tipen-demo/handler"

	"github.com/gofiber/fiber/v2"
)

type RegisterSellerBodyRequest struct {
	Storename string `json:"storename" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

type LoginSellerBodyRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateSellerBodyRequest struct {
	Storename string `json:"storename"`
}

func (r *Route) InitSeller() {
	api := r.Router.Group("/v1")
	seller := api.Group("/seller")
	seller.Post("/register", func(c *fiber.Ctx) error {
		p := RegisterSellerBodyRequest{}

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

		ID, err := r.Handler.RegisterSeller(handler.RegisterSellerParams{
			Storename: p.Storename,
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

	seller.Post("/login", func(c *fiber.Ctx) error {
		p := LoginSellerBodyRequest{}

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

		accessToken, err := r.Handler.LoginSeller(handler.LoginSellerParams{
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
	seller.Put("/update", func(c *fiber.Ctx) error {
		p := UpdateSellerBodyRequest{}

		sellerID, err := r.Middleware.AuthorizeSellerAndReturnID(c)
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
		if err := r.Handler.UpdateSeller(handler.UpdateSellerParams{
			ID:        sellerID,
			Storename: p.Storename,
		}); err != nil {
			return err
		}
		return c.SendString("Successfully Updated")
	})
	seller.Get("/profile", func(c *fiber.Ctx) error {
		sellerID, err := r.Middleware.AuthorizeSellerAndReturnID(c)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: err.Error(),
			}
		}
		resp, err := r.Handler.GetSellerProfile(sellerID)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrInternalServerError.Code,
				Message: err.Error(),
			}
		}
		return c.JSON(resp)
	})
}
