package _core

import (
	"fmt"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/digine/internal/_types"
)

func Bind(ptr any, labelPtr *string) {
	ptrType := _tools.ActualTypeOf(ptr)
	innerDMapper := (*_types.InnerDContainer)(nil)
	if ok := maputil.HasKey(container.TMap, ptrType); ok {
		innerDMapper = container.RequireSubContainer(ptrType)
	} else {
		innerDMapper = _types.NewInnerDIContainer()
		container.SetSubContainer(ptrType, innerDMapper)
	}
	// default label is address of pointer
	label := fmt.Sprint(ptr)
	if labelPtr != nil && !strutil.IsEmpty(*labelPtr) {
		label = *labelPtr
	}
	innerDMapper.SetValue(label, ptr)
}

func Inject(ptr any, labelPtr *string) {
	res := requireValue(_tools.ActualTypeOf(ptr), labelPtr)
	resVal := _tools.ActualValueOf(res)
	_tools.ActualValueOf(ptr).Set(resVal)
}
