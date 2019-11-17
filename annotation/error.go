package annotation

import (
	"fmt"
	"reflect"
	"strings"
)

type Error struct {
	message string
}

func NewNotNilError(name string) *Error {
	return &Error{message: fmt.Sprintf("Variable '%s' must be not nil", name)}
}

func NewNotEmptyError(name string) *Error {
	return &Error{message: fmt.Sprintf("Variable '%s' must be not empty", name)}
}

func NewErrorf(format string, arguments ...interface{}) error {
	return &Error{message: strings.TrimSpace(fmt.Sprintf(format, arguments...))}
}

func NewInvalidKindError(name string, value interface{}, expectedKinds ...reflect.Kind) *Error {
	parts := make([]string, len(expectedKinds))

	for i, expectedKind := range expectedKinds {
		parts[i] = expectedKind.String()
	}

	return &Error{
		message: fmt.Sprintf(
			"Variable '%s' (%T) must have one of kinds: %s",
			name,
			value,
			strings.Join(parts, ", "),
		),
	}
}

func (e *Error) Error() string {
	return e.message
}
