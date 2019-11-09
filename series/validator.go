package series

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// SizeIndex function to verify size value for series
func SizeIndex(
	v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string,
) bool {
	if size, ok := field.Interface().(int); ok {
		return size >= 0
	}

	return true
}
