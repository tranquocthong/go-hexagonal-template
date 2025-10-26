package domain

import "errors"

type ErrorCode string

const (
	ErrNotFound      ErrorCode = "not_found"
	ErrAlreadyExists ErrorCode = "already_exists"
	ErrInvalid       ErrorCode = "invalid"
	ErrUnauthorized  ErrorCode = "unauthorized"
	ErrForbidden     ErrorCode = "forbidden"
	ErrInternal      ErrorCode = "internal"
)

type DomainError struct {
	Code    ErrorCode
	Message string
}

func (e DomainError) Error() string { return e.Message }

func Is(err error, code ErrorCode) bool {
	var de DomainError
	if errors.As(err, &de) {
		return de.Code == code
	}
	return false
}
