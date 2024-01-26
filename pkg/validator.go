package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type Validator struct{}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var validate = validator.New()

func (v Validator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (v Validator) ErrorMessage(err ErrorResponse) string {
	if err.Tag == "required" {
		return fmt.Sprintf("%s wajib diisi", err.FailedField)
	} else if err.Tag == "min" {
		return fmt.Sprintf("%s minimal %s karakter", err.FailedField, err.Value)
	} else if err.Tag == "max" {
		return fmt.Sprintf("%s maksimal %s karakter", err.FailedField, err.Value)
	}
	return "unknown error"
}
