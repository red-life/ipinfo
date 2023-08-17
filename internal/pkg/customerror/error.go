package customerror

import "errors"

var (
	NotFoundErr     = errors.New("not found")
	NotSupportedErr = errors.New("not supported")
	GeneralErr      = errors.New("general error")
)
