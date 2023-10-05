package _core

import (
	"github.com/hanakogo/digine/internal/_types"
)

// WalkStructFields walk all fields of the struct,
// structure: struct object pointer
func WalkStructFields(stPtr any, fun func(option *_types.ReflectFieldOption)) {
	reflectOption := _types.NewReflectOption(stPtr)

	for _, reflectFieldOption := range reflectOption.Fields {
		fun(reflectFieldOption)
	}
}
