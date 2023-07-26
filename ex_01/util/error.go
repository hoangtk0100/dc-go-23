package util

import "errors"

var (
	ErrArgsFormatInvalid       = errors.New("input is not enough")
	ErrCountryCodeInvalid      = errors.New("country code is invalid")
	ErrCountryCodeNotSupported = errors.New("country code is not supported")
)
