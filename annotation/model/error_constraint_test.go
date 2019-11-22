package model

import (
	"fmt"

	"github.com/index0h/go-unit/unit"
	"github.com/pkg/errors"
)

type ErrorMessageConstraint struct {
	message string
}

func NewErrorMessageConstraint(format string, args ...interface{}) unit.Constraint {
	return &ErrorMessageConstraint{
		message: fmt.Sprintf(format, args...),
	}
}

func (c *ErrorMessageConstraint) Check(expected interface{}) bool {
	err, ok := expected.(error)

	if !ok {
		panic(errors.Errorf("Variable 'expected' (%T) must have type: error", expected))
	}

	return err.Error() == c.message
}

func (c *ErrorMessageConstraint) String() string {
	return fmt.Sprintf("have equal message to:\n%s", c.message)
}

func (c *ErrorMessageConstraint) Details(actual interface{}) string {
	return fmt.Sprintf("actual:\n%s", actual.(error).Error())
}
