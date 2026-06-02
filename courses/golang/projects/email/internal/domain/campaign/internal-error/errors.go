package internalerror

import "errors"

var (
	ErrInternalError error = errors.New("Internal Server Error")
)
