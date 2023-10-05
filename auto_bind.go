package digine

import (
	"github.com/hanakogo/digine/internal/_core"
	"github.com/hanakogo/digine/internal/_tools"
)

// AutoBindFields inject fields with digine tag
func AutoBindFields(obj any) {
	if !_tools.IsPointer(obj) {
		obj = &obj
	}
	_core.AutoBindFields(obj)
}
