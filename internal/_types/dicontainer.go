package _types

import (
	"github.com/gookit/goutil/maputil"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/hanakoutilgo"
	"reflect"
)

// TODO: match interface type to inject

// base types
type (
	TMap[T comparable, R any]       map[T]R
	DContainer[T comparable, R any] struct {
		TMap[T, R]
	}
)

type (
	OuterDContainer DContainer[reflect.Type, *InnerDContainer]
	InnerDContainer DContainer[string, any]
)

func NewOuterDIContainer() *OuterDContainer {
	return &OuterDContainer{
		map[reflect.Type]*InnerDContainer{},
	}
}

func NewInnerDIContainer() *InnerDContainer {
	return &InnerDContainer{
		map[string]any{},
	}
}

func (d *OuterDContainer) RequireSubContainer(key reflect.Type) *InnerDContainer {
	var dContainer *InnerDContainer = nil
	if ok := maputil.HasKey(d.TMap, key); ok {
		res := d.TMap[key]
		hanakoutilgo.CastThen[*InnerDContainer](res, func(container *InnerDContainer) {
			dContainer = container
		})
	}
	return dContainer
}

func (d *OuterDContainer) SetSubContainer(key reflect.Type, container *InnerDContainer) {
	d.TMap[key] = container
}

func (d *InnerDContainer) RequireValue(key string, types reflect.Type) any {
	var res any = nil
	if ok := maputil.HasKey(d.TMap, key); ok {
		res = d.TMap[key]
	}
	if res != nil && _tools.ActualTypeOf(res) == types {
		return res
	}
	return nil
}

func (d *InnerDContainer) SetValue(key string, value any) {
	d.TMap[key] = value
}
