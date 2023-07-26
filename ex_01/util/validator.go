package util

import (
	"strings"
)

func ValidateArgs(input []string) error {
	length := len(input)
	if length < MinArgsLength {
		return ErrArgsFormatInvalid
	}

	return ValidateCountryCode(input[length-1])
}

func ValidateCountryCode(input string) error {
	input = strings.TrimSpace(input)
	if len(input) < MinCountryCodeLength {
		return ErrCountryCodeInvalid
	}

	switch CountryCode(strings.ToUpper(input)) {
	case US, VN:
		return nil
	default:
		return ErrCountryCodeNotSupported
	}
}
