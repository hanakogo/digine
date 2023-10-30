package test

import (
	"github.com/hanakogo/digine"
	"github.com/hanakogo/digine/internal/_tools"
	"github.com/hanakogo/exceptiongo"
	"testing"
)

func TestStructResolve(t *testing.T) {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		message := exception.GetStackTraceMessage()
		t.Error(message)
	}).Deploy()
	BindStruct(t)
	InjectStruct(t)
}

type StructToInjWrapper struct {
	Struct *StructToInj `digine:"struct1"`
}

type StructToInj struct {
	StrD string `digine:"strD"`
}

func BindStruct(t *testing.T) {
	struct1 := &StructToInj{"struct1"}
	digine.Bind[StructToInj](struct1, digine.NewLabel("struct1"))
	digine.StructBindFields(struct1)
	t.Log(struct1)
}

func InjectStruct(t *testing.T) {
	reqStruct1 := digine.StructRequire[StructToInjWrapper]()
	t.Log(reqStruct1.Struct)
	var struct1 = _tools.NewValuePtr[StructToInj]()
	digine.StructInject[StructToInj](struct1)
	t.Log(struct1)
}
