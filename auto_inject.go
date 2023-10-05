package digine

import (
	"github.com/hanakogo/digine/internal/_core"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/digine/internal/_types"
	"github.com/hanakogo/exceptiongo"
)

// AutoInject formally extract dependencies
func AutoInject[T any](ptr *T) {
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

// AutoRequire a dirty way (for GC) to extract dependencies with a new object
func AutoRequire[T any]() *T {
	ptr := _tools.NewValuePtr[T]()
	AutoInject[T](ptr)
	return ptr
}
