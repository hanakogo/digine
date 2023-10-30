package digine

import (
	"github.com/hanakogo/digine/internal/_core"
	"github.com/hanakogo/digine/internal/_tools"
)

// StructBindFields inject fields with digine tag
func StructBindFields(obj any) {
	if !_tools.IsPointer(obj) {
		obj = &obj
	}
	_core.AutoBindFields(obj)
}
