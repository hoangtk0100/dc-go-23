package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
	. "github.com/hoangtk0100/dc-go-23/ex_08/pkg/constant"
)

var (
	ErrPasswordTooShort = errors.New("password too short, password must be at least 8 characters")
)

func IsCurrency(input string) bool {
	currency := Currency(input)
	switch currency {
	case CurrencyUSD, CurrencyEUR, CurrencyVND:
		return true
	default:
		return false
	}
}

func IsWeightUnit(input string) bool {
	unit := WeightUnit(input)
	switch unit {
	case WeightUnitGram, WeightUnitLBS, WeightUnitKg:
		return true
	default:
		return false
	}
}

var ValidateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return IsCurrency(currency)
	}

	return false
}

var ValidateWeightUnit validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if unit, ok := fieldLevel.Field().Interface().(string); ok {
		return IsWeightUnit(unit)
	}

	return false
}

func IsProductStatus(input string) bool {
	unit := ProductStatus(input)
	switch unit {
	case ProductStatusActive, ProductStatusDeleted:
		return true
	default:
		return false
	}
}

func ValidatePassword(input string) error {
	if len(input) < 8 {
		return ErrPasswordTooShort
	}

	return nil
}
