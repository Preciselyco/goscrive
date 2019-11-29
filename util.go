package scrive

import (
	"reflect"
)

func anyNil(values ...interface{}) bool {
	for _, v := range values {
		if v == nil {
			return true
		}
		vVal := reflect.ValueOf(v)
		if vVal.Kind() == reflect.Ptr && vVal.IsNil() {
			return true
		}
	}
	return false
}

func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
