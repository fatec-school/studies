package internalerror

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s any) error {
	validate := validator.New()
	err := validate.Struct(s)

	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.Field() + " is required")
	case "email":
		return errors.New(validationError.Field() + " must be a valid email")
	case "min":
		return errors.New(validationError.Field() + " must be at least " + validationError.Param() + " characters long")
	case "max":
		return errors.New(validationError.Field() + " must be at most " + validationError.Param() + " characters long")
	default:
		return ErrValidationError
	}

}
