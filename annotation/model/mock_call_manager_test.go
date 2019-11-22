package model

import "github.com/index0h/go-unit/unit"

type MockCallManager struct {
	ctrl    *unit.Controller
	calls   []*MockCall
	options []interface{}
}

func NewMockCallManager(ctrl *unit.Controller, options ...interface{}) *MockCallManager {
	result := &MockCallManager{
		ctrl:    ctrl,
		calls:   []*MockCall{},
		options: options,
	}

	ctrl.RegisterFinish(result.Finish)

	return result
}

func (m *MockCallManager) CreateCall(methodName string, arguments ...interface{}) *MockCall {
	m.ctrl.TestingT().Helper()

	call := NewMockCall(methodName, unit.NewEqualComparator(), arguments...)

	m.calls = append(m.calls, call)

	return call
}

func (m *MockCallManager) FetchCall(methodName string, arguments ...interface{}) *MockCall {
	m.ctrl.TestingT().Helper()

	for i, call := range m.calls {
		if call != nil {
			if !call.arguments.Check(arguments) {
				m.ctrl.TestingT().Errorf(
					"Failed assertion for method call %s arguments.\n%s",
					methodName,
					call.arguments.Details(arguments),
				)
			}

			m.calls[i] = nil

			return call
		}
	}

	m.ctrl.TestingT().Errorf("Mock call for method '%s' is not registered", methodName)
	m.ctrl.TestingT().Fail()

	return nil
}

func (m *MockCallManager) Finish() {
	m.ctrl.TestingT().Helper()

	for _, call := range m.calls {
		if call != nil {
			m.ctrl.TestingT().Errorf("Mock call for method '%s' is registered, but never called", call.methodName)
		}
	}
}
