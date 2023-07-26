package validator

import (
	"errors"
	"github.com/hoangtk0100/dc-go-23/ex_02/constant"
	"strconv"
)

var (
	ErrInputInvalidType = errors.New("input values are not match type")
)

func ValidateInputs(inputType constant.InputType, inputs []string) ([]interface{}, error) {
	result := make([]interface{}, len(inputs))

	switch inputType {
	case constant.IntInputType:
		for index, val := range inputs {
			parsedValue, err := strconv.Atoi(val)
			if err != nil {
				return nil, ErrInputInvalidType
			}

			result[index] = parsedValue
		}

	case constant.FloatInputType:
		for index, val := range inputs {
			parsedValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, ErrInputInvalidType
			}

			result[index] = parsedValue
		}

	case constant.StringInputType:
		for index, val := range inputs {
			result[index] = val
		}

	case constant.MixInputType:
		for index, val := range inputs {
			if parsedValue, err := strconv.Atoi(val); err == nil {
				result[index] = parsedValue
				continue
			}

			if parsedValue, err := strconv.ParseFloat(val, 64); err == nil {
				result[index] = parsedValue
				continue
			}

			result[index] = val
		}
	}

	return result, nil
}
