package util

import (
	"fmt"
	"strings"
)

type NameFormatter interface {
	Format(firstName, lastName, middleName string) string
}

func NewNameFormatter(countryCode string) NameFormatter {
	switch CountryCode(strings.ToUpper(countryCode)) {
	case VN:
		return &vnFormatter{}
	case US:
		return &usFormatter{}
	default:
		return &usFormatter{}
	}
}

type (
	vnFormatter struct{}
	usFormatter struct{}
)

func (f *vnFormatter) Format(firstName, lastName, middleName string) string {
	return fmt.Sprintf("%s%s %s", lastName, getFormattedMiddleName(middleName), firstName)
}

func (f *usFormatter) Format(firstName, lastName, middleName string) string {
	return fmt.Sprintf("%s %s%s", firstName, lastName, getFormattedMiddleName(middleName))
}

func getFormattedMiddleName(input string) string {
	if len(input) > 0 {
		return " " + input
	}

	return input
}
