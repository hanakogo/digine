package test

import (
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"testing"
)

func TestBaseDI(t *testing.T) {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		message := exception.GetStackTraceMessage()
		t.Error(message)
	}).Deploy()
	Inject(t)
	Extract(t)
}

func Inject(t *testing.T) {
	var strA = "test string"
	var numA = 1919
	digine.Bind[string](&strA, digine.NewLabel("strA"))
	digine.Bind[int](&numA, digine.NewLabel("numA"))
}

func Extract(t *testing.T) {
	strA := digine.Require[string](digine.NewLabel("strA"))
	numA := digine.Require[int](digine.NewLabel("numA"))
	t.Log(*strA, *numA)
}
