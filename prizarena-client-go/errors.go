package prizarena

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorized = errors.New("unauthorized");
	ErrForbidden    = errors.New("forbidden")
)

type apiError struct {
	Code string
	Message string
}

func (err apiError) Error() string {
	return fmt.Sprintf("%v: %v", err.Code, err.Message)
}

type ErrBadRequest struct {
	apiError
}

type ErrInternalServerError struct {
	apiError
}
