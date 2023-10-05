package _core

import (
	"fmt"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/digine/internal/_types"
	"log/slog"
	"reflect"
)

func AutoInject(ptr any) {
	WalkStructFields(ptr, func(option *_types.ReflectFieldOption) {
		if !option.IsExported() {
			slog.Warn(fmt.Sprintf("skip un-exported typeField: %s", option.Field.Name))
			return
		}
		if option.IsStruct && option.DigineTag == "@" {
			// if this field is a nil pointer, we must initialize it to AutoInject
			if option.IsNil() {
				_tools.InjectNewValuePtr(option.FieldType, option.FieldValue)
			}
			AutoInject(option.ValuePtr())
		} else {
			injectValuePtr := requireValue(option.FieldType, &option.DigineTag)
			if option.IsPtrField {
				option.FieldValue.Set(reflect.ValueOf(injectValuePtr))
			} else {
				option.FieldValue.Set(reflect.ValueOf(injectValuePtr).Elem())
			}
		}
	})
}
