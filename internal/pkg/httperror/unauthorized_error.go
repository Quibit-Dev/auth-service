package httperror

import (
	"errors"
	"net/http"

	"auth-service/internal/pkg/constant"
)

func NewUnauthorizedError() *ResponseError {
	msg := constant.UnauthorizedErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusUnauthorized, msg)
}
