package internalerror

import "errors"

var (
	ErrInternalError error = errors.New("Internal Server Error")
	ErrValidationError error = errors.New("Validation Error")
)
