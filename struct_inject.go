package digine

import (
	"github.com/hanakogo/digine/internal/_core"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/digine/internal/_types"
	"github.com/hanakogo/exceptiongo"
)

// StructInject formally extract fields into struct
func StructInject[T any](ptr *T) {
	if !_tools.IsStructPointer(ptr) {
		e := exceptiongo.NewException[_types.NotSupportedException]("only struct type is supported")
		exceptiongo.Throw(e)
	}
	if ptr == nil {
		e := exceptiongo.NewException[_types.NullPointerException]("can't inject nil pointer")
		exceptiongo.Throw(e)
	}
	_core.AutoInject(ptr)
}

// StructRequire a dirty way (for GC) to extract dependencies with a struct pointer
func StructRequire[T any]() *T {
	ptr := _tools.NewValuePtr[T]()
	StructInject[T](ptr)
	return ptr
}
