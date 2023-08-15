package util

import (
	"net/http"
)

type DefaultError struct {
	StatusField string `json:"status,omitempty"`
	CodeField   int    `json:"code,omitempty"`
	ErrorField  string `json:"message"`
	DebugField  string `json:"debug,omitempty"`
}

func (e DefaultError) WithError(message string) *DefaultError {
	e.ErrorField = message
	return &e
}

func (e DefaultError) WithDebug(debug string) *DefaultError {
	e.DebugField = debug
	return &e
}

func (e DefaultError) Error() string {
	return e.ErrorField
}

func (e DefaultError) StatusCode() int {
	return e.CodeField
}

func (e DefaultError) Is(err error) bool {
	switch te := err.(type) {
	case DefaultError:
		return e.ErrorField == te.ErrorField &&
			e.StatusField == te.StatusField &&
			e.CodeField == te.CodeField
	case *DefaultError:
		return e.ErrorField == te.ErrorField &&
			e.StatusField == te.StatusField &&
			e.CodeField == te.CodeField
	default:
		return false
	}
}

type StatusCodeCarrier interface {
	StatusCode() int
}

var ErrInternalServerError = DefaultError{
	StatusField: http.StatusText(http.StatusInternalServerError),
	ErrorField:  "An internal server error occurred. Please try again later",
	CodeField:   http.StatusInternalServerError,
}

var ErrUnauthorized = DefaultError{
	StatusField: http.StatusText(http.StatusUnauthorized),
	ErrorField:  "Access denied. Please authenticate with valid credentials",
	CodeField:   http.StatusUnauthorized,
}

var ErrBadRequest = DefaultError{
	StatusField: http.StatusText(http.StatusBadRequest),
	ErrorField:  "The request was invalid or contained malformed parameters",
	CodeField:   http.StatusBadRequest,
}

var ErrNotFound = DefaultError{
	StatusField: http.StatusText(http.StatusNotFound),
	ErrorField:  "The requested page or resource could not be found",
	CodeField:   http.StatusNotFound,
}

var ErrForbidden = DefaultError{
	StatusField: http.StatusText(http.StatusForbidden),
	ErrorField:  "Access to the requested page or resource is forbidden",
	CodeField:   http.StatusForbidden,
}

var ErrUnsupportedMediaType = DefaultError{
	StatusField: http.StatusText(http.StatusUnsupportedMediaType),
	ErrorField:  "The media type of the requested resource is not supported",
	CodeField:   http.StatusUnsupportedMediaType,
}

var ErrConflict = DefaultError{
	StatusField: http.StatusText(http.StatusConflict),
	ErrorField:  "A conflict occurred with the current state of the resource",
	CodeField:   http.StatusConflict,
}

var ErrTimeout = DefaultError{
	StatusField: http.StatusText(http.StatusRequestTimeout),
	ErrorField:  "The request timed out",
	CodeField:   http.StatusRequestTimeout,
}
