package route

import (
	"strings"
	"tipen-demo/handler"

	"github.com/gofiber/fiber/v2"
)

type CreateProductBodyRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Image string `json:"image" validate:"required"`
}

type UpdateProductBodyRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (r *Route) InitProduct() {
	api := r.Router.Group("/v1")
	product := api.Group("/products")
	product.Post("", func(c *fiber.Ctx) error {
		p := CreateProductBodyRequest{}

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

		ID, err := r.Handler.CreateProduct(handler.CreateProductParams{
			Name:     p.Name,
			Price:    p.Price,
			Image:    p.Image,
			SellerID: sellerID,
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

	product.Put("/:id", func(c *fiber.Ctx) error {
		p := UpdateProductBodyRequest{}

		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		_, err = r.Middleware.AuthorizeSellerAndReturnID(c)
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
		if err := r.Handler.UpdateProduct(handler.UpdateProductParams{
			ID:    id,
			Name:  p.Name,
			Price: p.Price,
			Image: p.Image,
		}); err != nil {
			return err
		}
		return c.SendString("Successfully Updated")
	})
	product.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		resp, err := r.Handler.GetProductByID(id)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrInternalServerError.Code,
				Message: err.Error(),
			}
		}
		return c.JSON(resp)
	})
	// product.Get("", func(c *fiber.Ctx) error {

	// })
}
