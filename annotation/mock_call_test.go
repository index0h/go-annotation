package annotation

import "github.com/index0h/go-unit/unit"

const (
	MockCallTypeNotConfigured int8 = iota
	MockCallTypeReturn
	MockCallTypePanic
	MockCallTypeCallback
)

type MockCall struct {
	methodName   string
	arguments    *unit.ElementsConstraint
	mockCallType int8
	returns      []interface{}
	panics       interface{}
	callback     interface{}
}

func NewMockCall(methodName string, comparator unit.EqualComparer, arguments ...interface{}) *MockCall {
	return &MockCall{
		methodName: methodName,
		arguments:  unit.NewValueElementsConstraint(comparator, arguments...).(*unit.ElementsConstraint),
	}
}

func (c *MockCall) SetReturn(values ...interface{}) {
	if c.mockCallType != MockCallTypeNotConfigured {
		panic(unit.NewErrorf("Call for row %s already configured", c.methodName))
	}

	c.mockCallType = MockCallTypeReturn
	c.returns = values
}

func (c *MockCall) SetPanic(value interface{}) {
	if c.mockCallType != MockCallTypeNotConfigured {
		panic(unit.NewErrorf("Call for row %s already configured", c.methodName))
	}

	c.mockCallType = MockCallTypePanic
	c.panics = value
}

func (c *MockCall) SetCallback(callback interface{}) {
	if c.mockCallType != MockCallTypeNotConfigured {
		panic(unit.NewErrorf("Call for row %s already configured", c.methodName))
	}

	c.mockCallType = MockCallTypeCallback
	c.callback = callback
}

func (c *MockCall) Call() (interface{}, int8) {
	switch c.mockCallType {
	case MockCallTypeReturn:
		return c.returns, c.mockCallType
	case MockCallTypePanic:
		return c.panics, c.mockCallType
	case MockCallTypeCallback:
		return c.callback, c.mockCallType
	default:
		panic(unit.NewErrorf("Call for row %s is not configured", c.methodName))
	}
}
