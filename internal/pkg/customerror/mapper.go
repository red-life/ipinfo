package customerror

import (
	"errors"
	"net/http"
)

func MaxMindToCustom(err error) error {
	switch {
	default:
		return GeneralErr
	}
}

func ErrorToStatusCode(err error) int {
	switch {
	case errors.Is(err, NotFoundErr):
		return http.StatusNotFound
	case errors.Is(err, NotSupportedErr):
		return http.StatusNotImplemented
	case errors.Is(err, GeneralErr):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
