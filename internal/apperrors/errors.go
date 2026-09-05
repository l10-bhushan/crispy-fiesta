package apperrors

import "errors"

var (
	ErrorNotFound     = errors.New("resource not found")
	ErrorInternal     = errors.New("internal server error")
	ErrorInvalidInput = errors.New("invalid input")
	ErrorConflict     = errors.New("resource already exists")
)
