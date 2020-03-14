package common

import (
	"fmt"
	"reflect"
)

func IsValidInterface(data interface{}, kind reflect.Kind, name string) (bool, error) {
	if data == nil {
		return false, fmt.Errorf("invalid param %s should not be nil", name)
	}

	if reflect.TypeOf(data).Kind() == kind {
		return false, fmt.Errorf("invalid param %s should be %s but %s", name, kind.String(), reflect.TypeOf(data).Kind().String())
	}

	return true, nil
}
