package _tools

import (
	"reflect"
)

func IsPointer(o any) bool {
	types := reflect.TypeOf(o)
	if types.Kind() == reflect.Ptr {
		return true
	}
	return false
}

func IsStructPointer(o any) bool {
	if !IsPointer(o) {
		return false
	}
	kind := reflect.TypeOf(o).Elem().Kind()
	if kind == reflect.Struct {
		return true
	}
	return false
}

func ActualTypeOf(o any) reflect.Type {
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func ActualValueOf(o any) reflect.Value {
	t := reflect.ValueOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
