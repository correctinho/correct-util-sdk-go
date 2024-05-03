package utilsdk

import (
	"reflect"
)

// IsNil - valida se eh nulo
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	if reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil() {
		return true
	}
	if reflect.ValueOf(value).Kind() == reflect.Invalid {
		return true
	}
	return false
}
