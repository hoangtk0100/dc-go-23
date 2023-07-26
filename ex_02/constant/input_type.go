package constant

import "errors"

var (
	ErrInputTypeMissing = errors.New("missing input type")
)

type InputType string

const (
	IntInputType    InputType = "int"
	FloatInputType  InputType = "float"
	StringInputType InputType = "string"
	MixInputType    InputType = "mix"
)

func GetInputType(intFlag, floatFlag, stringFlag, mixFlag *bool) (InputType, error) {
	if *intFlag {
		return IntInputType, nil
	}

	if *floatFlag {
		return FloatInputType, nil
	}

	if *stringFlag {
		return StringInputType, nil
	}

	if *mixFlag {
		return MixInputType, nil
	}

	return "", ErrInputTypeMissing
}
