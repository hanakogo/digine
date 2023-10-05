package _types

import (
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/exceptiongo"
	"reflect"
)

type ReflectOption struct {
	FieldCount int
	IsPtr      bool

	ObjType  reflect.Type
	ObjValue reflect.Value
	Fields   []*ReflectFieldOption
}

func NewReflectOption(ptr any) *ReflectOption {
	if !_tools.IsPointer(ptr) {
		e := exceptiongo.NewException[NotSupportedException]("only supported for pointer")
		exceptiongo.Throw(e)
	}

	typ := _tools.ActualTypeOf(ptr)
	value := _tools.ActualValueOf(ptr)

	if typ.Kind() != reflect.Struct {
		e := exceptiongo.NewException[NotSupportedException]("not supported unless struct")
		exceptiongo.Throw(e)
	}

	fieldOptions := make([]*ReflectFieldOption, 0)
	numField := typ.NumField()
	for i := 0; i < numField; i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)
		fieldType := field.Type
		isPtrField := fieldValue.Kind() == reflect.Ptr

		if isPtrField {
			fieldType = fieldType.Elem()
		}

		digineTag := field.Tag.Get("digine")
		if digineTag == "" {
			continue
		}

		fieldOptions = append(fieldOptions, &ReflectFieldOption{
			IsPtrField: isPtrField,
			IsStruct:   fieldType.Kind() == reflect.Struct,
			Field:      field,
			FieldType:  fieldType,
			FieldValue: fieldValue,
			DigineTag:  digineTag,
		})
	}
	return &ReflectOption{
		FieldCount: numField,
		IsPtr:      _tools.IsPointer(ptr),
		ObjType:    typ,
		ObjValue:   value,
		Fields:     fieldOptions,
	}
}

type ReflectFieldOption struct {
	DigineTag  string
	IsPtrField bool
	IsStruct   bool

	Field      reflect.StructField
	FieldType  reflect.Type
	FieldValue reflect.Value
}

// ValuePtr get pointer of actual value
func (r *ReflectFieldOption) ValuePtr() any {
	if r.IsPtrField && !r.FieldValue.IsNil() {
		return r.FieldValue.Addr().Elem().Interface()
	}
	return r.FieldValue.Addr().Interface()
}

func (r *ReflectFieldOption) IsNil() bool {
	if r.IsPtrField {
		return r.FieldValue.IsNil()
	}
	return false
}

func (r *ReflectFieldOption) IsExported() bool {
	return r.Field.IsExported()
}

func (r *ReflectFieldOption) New() {
	r.FieldValue.Set(reflect.New(r.FieldType))
}
