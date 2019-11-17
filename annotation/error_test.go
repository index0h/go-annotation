package annotation

import (
	"github.com/index0h/go-unit/unit"
	"reflect"
	"testing"
)

func TestNewNotNilError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewNotNilError, "variable").
		ExpectResult(&Error{message: "Variable 'variable' must be not nil"})
}

func TestNewNotEmptyError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewNotEmptyError, "variable").
		ExpectResult(&Error{message: "Variable 'variable' must be not empty"})
}

func TestNewErrorf(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewErrorf, "data %T", [1]int{5}).
		ExpectResult(&Error{message: "data [1]int"})
}

func TestNewInvalidKindError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewInvalidKindError, "variable", "data", reflect.Int, reflect.Bool).
		ExpectResult(&Error{message: "Variable 'variable' (string) must have one of kinds: int, bool"})
}

func TestError_Error(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call((&Error{message: "data"}).Error).
		ExpectResult("data")
}
