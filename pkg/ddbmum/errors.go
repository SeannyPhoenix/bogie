package ddbmum

import (
	"fmt"
	"reflect"
)

type UnsupportedTypeError struct {
	Kind reflect.Kind
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type: %s", e.Kind)
}
