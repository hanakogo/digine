package _core

import (
	"fmt"
	"github.com/hanakogo/digine/internal/_types"
	"log/slog"
)

func AutoBindFields(ptr any) {
	WalkStructFields(ptr, func(option *_types.ReflectFieldOption) {
		if option.IsNil() {
			return
		}
		if !option.IsExported() {
			slog.Warn(fmt.Sprintf("digine/AutoBindFields: skip un-exported typeField: %s", option.Field.Name))
			return
		}
		if option.IsStruct && option.DigineTag == "@" {
			AutoBindFields(option.ValuePtr())
		} else {
			if !option.IsPtrField {
				slog.Warn("digine/AutoBindFields: inject a basic value with non-pointer type may cause unexpected behavior, please use pointer type instead for better")
			}
			Bind(option.ValuePtr(), &option.DigineTag)
		}
	})
}
