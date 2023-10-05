package test

import (
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"testing"
)

func TestAutoBindInject(t *testing.T) {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		message := exception.GetStackTraceMessage()
		t.Error(message)
	}).Deploy()

	AutoBind(t)
	AutoInject(t)
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

func AutoBind(t *testing.T) {
	strA := "abc"
	strB := "def"
	strC := "ghi"
	digine.AutoBindFields(&St{
		StrA:  &strA,
		StrB:  strB,
		StEl2: StEl{StrC: &strC},
	})
}

func AutoInject(t *testing.T) {
	ptrTest := &St{}
	digine.AutoInject[St](ptrTest)
	t.Log(*ptrTest)

	require := digine.AutoRequire[St]()
	t.Log(*require)
}
