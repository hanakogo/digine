package _core

import (
	"fmt"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/hanakogo/digine/internal/_types"
	"github.com/hanakogo/exceptiongo"
	"log/slog"
	"reflect"
)

var container = _types.NewOuterDIContainer()

func requireValue(ptrType reflect.Type, labelPtr *string) (val any) {
	defer func() {
		if val != nil {
			return
		}
		label := "nil"
		if labelPtr != nil {
			label = *labelPtr
		}
		f := exceptiongo.NewExceptionF[_types.NullPointerException]("can't found the object that labeled by %s", label)
		exceptiongo.Throw(f)
	}()

	var innerDMapper *_types.InnerDContainer
	if ok := maputil.HasKey(container.TMap, ptrType); ok {
		innerDMapper = container.RequireSubContainer(ptrType)
	}
	if innerDMapper == nil {
		return nil
	}

	if labelPtr != nil && !strutil.IsEmpty(*labelPtr) {
		val = innerDMapper.RequireValue(*labelPtr, ptrType)
		return val
	}

	keys := maputil.Keys(innerDMapper.TMap)
	if len(keys) >= 1 {
		if len(keys) > 1 {
			slog.Warn(fmt.Sprintf("digine/_core/dep_inj: found %d object of type %v, extract object will be inaccurate", len(keys), ptrType))
		}
		val = innerDMapper.RequireValue(keys[0], ptrType)
		return val
	}

	return
}
