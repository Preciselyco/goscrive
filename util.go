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

func Bool(value bool) *bool {
	return &value
}

func String(value string) *string {
	return &value
}

func UInt(value uint64) *uint64 {
	return &value
}

func UInt32(value uint32) *uint32 {
	return &value
}
