package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(errs validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, err := range errs {
		var message string

		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", err.Field())
		case "gte":
			message = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
		case "lte":
			message = fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
		default:
			message = fmt.Sprintf("%s is not valid", err.Field())

		}
		errors[err.Field()] = message
	}
	return errors
}