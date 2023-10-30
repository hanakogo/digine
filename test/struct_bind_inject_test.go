package test

import (
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"testing"
)

func TestStructBindInject(t *testing.T) {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		message := exception.GetStackTraceMessage()
		t.Error(message)
	}).Deploy()

	StructBind(t)
	StructInject(t)
}

type St struct {
	StrA  *string `digine:"A"`
	StrB  string  `digine:"B"`
	StEl1 *StEl   `digine:"@"`
	StEl2 StEl    `digine:"@"`
}

type StEl struct {
	StrC *string `digine:"C"`
}

func StructBind(t *testing.T) {
	strA := "abc"
	strB := "def"
	strC := "ghi"
	digine.StructBindFields(&St{
		StrA:  &strA,
		StrB:  strB,
		StEl2: StEl{StrC: &strC},
	})
}

func StructInject(t *testing.T) {
	ptrTest := &St{}
	digine.StructInject[St](ptrTest)
	t.Log(*ptrTest)

	require := digine.StructRequire[St]()
	t.Log(*require)
}
