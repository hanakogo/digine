package digine

import (
	"github.com/hanakogo/digine/internal/_core"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/digine/internal/_types"
	"github.com/hanakogo/exceptiongo"
	"log/slog"
)

func Bind[T any](obj *T, label *Label) {
	if obj == nil {
		slog.Warn("digine/di: can't inject nil object")
		return
	}
	_core.Bind(obj, label.Get())
}

func Inject[T any](ptr *T, label *Label) {
	if ptr == nil {
		e := exceptiongo.NewException[_types.NullPointerException]("can't inject nil pointer")
		exceptiongo.Throw(e)
	}

	// for special "@" tag
	labelPtr := label.Get()
	if labelPtr != nil && *labelPtr == "@" {
		labelPtr = nil
	}

	_core.Inject(ptr, labelPtr)
}

func Require[T any](label *Label) *T {
	ptr := _tools.NewValuePtr[T]()
	Inject[T](ptr, label)
	return ptr
}
