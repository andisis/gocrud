package helper

import "net/http"

// GetStatusCode represent http status code
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadParamInput:
		return http.StatusBadRequest
	case ErrMissingParam:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
