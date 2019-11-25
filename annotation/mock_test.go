package annotation

import (
	"github.com/index0h/go-unit/unit"
	"github.com/pkg/errors"
)

type (
	SpecMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	SpecMockRecorder struct {
		mock *SpecMock
	}

	SpecMockRecorderForString struct {
		call *MockCall
	}

	SpecMockRecorderForValidate struct {
		call *MockCall
	}

	SpecMockRecorderForClone struct {
		call *MockCall
	}

	SpecMockRecorderForFetchImports struct {
		call *MockCall
	}

	SpecMockRecorderForRenameImports struct {
		call *MockCall
	}

	SpecMockRecorderForIsEqual struct {
		call *MockCall
	}
)

type (
	AnnotationParserMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	AnnotationParserMockRecorder struct {
		mock *AnnotationParserMock
	}

	AnnotationParserMockRecorderForParse struct {
		call *MockCall
	}
)

type (
	SourceParserMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	SourceParserMockRecorder struct {
		mock *SourceParserMock
	}

	SourceParserMockRecorderForParse struct {
		call *MockCall
	}
)

type (
	MarshalerJSONMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	MarshalerJSONMockRecorder struct {
		mock *MarshalerJSONMock
	}

	MarshalerJSONMockRecorderForMarshalJSON struct {
		call *MockCall
	}
)

func NewSpecMock(ctrl *unit.Controller, options ...interface{}) *SpecMock {
	return &SpecMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *SpecMock) EXPECT() *SpecMockRecorder {
	return &SpecMockRecorder{mock: m}
}

func (m *SpecMock) String() (result string) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("String").Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(string)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() (result string))()
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) String() *SpecMockRecorderForString {
	mr.mock.ctrl.TestingT().Helper()

	return &SpecMockRecorderForString{
		call: mr.mock.callManager.CreateCall("String"),
	}
}

func (mrm *SpecMockRecorderForString) Return(result string) {
	mrm.call.SetReturn(result)
}

func (mrm *SpecMockRecorderForString) String(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForString) Callback(callback func() (result string)) {
	mrm.call.SetCallback(callback)
}

func (m *SpecMock) Validate() {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("Validate").Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func())()
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) Validate() *SpecMockRecorderForValidate {
	mr.mock.ctrl.TestingT().Helper()

	return &SpecMockRecorderForValidate{
		call: mr.mock.callManager.CreateCall("Validate"),
	}
}

func (mrm *SpecMockRecorderForValidate) Return() {
	mrm.call.SetReturn()
}

func (mrm *SpecMockRecorderForValidate) Validate(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForValidate) Callback(callback func()) {
	mrm.call.SetCallback(callback)
}

func (m *SpecMock) Clone() (result interface{}) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("Clone").Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(interface{})
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() (result interface{}))()
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) Clone() *SpecMockRecorderForClone {
	mr.mock.ctrl.TestingT().Helper()

	return &SpecMockRecorderForClone{
		call: mr.mock.callManager.CreateCall("Clone"),
	}
}

func (mrm *SpecMockRecorderForClone) Return(result interface{}) {
	mrm.call.SetReturn(result)
}

func (mrm *SpecMockRecorderForClone) Clone(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForClone) Callback(callback func() (result interface{})) {
	mrm.call.SetCallback(callback)
}

func (m *SpecMock) FetchImports(file *File) (result0 []*Import) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, file)

	switch __result, __type := m.callManager.FetchCall("FetchImports", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].([]*Import)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(file *File) []*Import)(file)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) FetchImports(file interface{}) *SpecMockRecorderForFetchImports {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, file)

	return &SpecMockRecorderForFetchImports{
		call: mr.mock.callManager.CreateCall("FetchImports", __params...),
	}
}

func (mrm *SpecMockRecorderForFetchImports) Return(result0 []*Import) {
	mrm.call.SetReturn(result0)
}

func (mrm *SpecMockRecorderForFetchImports) FetchImports(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForFetchImports) Callback(callback func(file *File) []*Import) {
	mrm.call.SetCallback(callback)
}

func (m *SpecMock) RenameImports(oldAlias string, newAlias string) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, oldAlias)
	__params = append(__params, newAlias)

	switch __result, __type := m.callManager.FetchCall("RenameImports", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(oldAlias string, newAlias string))(oldAlias, newAlias)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) RenameImports(oldAlias interface{}, newAlias interface{}) *SpecMockRecorderForRenameImports {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, oldAlias)
	__params = append(__params, newAlias)

	return &SpecMockRecorderForRenameImports{
		call: mr.mock.callManager.CreateCall("RenameImports", __params...),
	}
}

func (mrm *SpecMockRecorderForRenameImports) Return() {
	mrm.call.SetReturn()
}

func (mrm *SpecMockRecorderForRenameImports) RenameImports(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForRenameImports) Callback(callback func(oldAlias string, newAlias string)) {
	mrm.call.SetCallback(callback)
}

func (m *SpecMock) IsEqual(value interface{}) (result bool) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, value)

	switch __result, __type := m.callManager.FetchCall("IsEqual", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(bool)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(value interface{}) (result bool))(value)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SpecMockRecorder) IsEqual(value interface{}) *SpecMockRecorderForIsEqual {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, value)

	return &SpecMockRecorderForIsEqual{
		call: mr.mock.callManager.CreateCall("IsEqual", __params...),
	}
}

func (mrm *SpecMockRecorderForIsEqual) Return(result bool) {
	mrm.call.SetReturn(result)
}

func (mrm *SpecMockRecorderForIsEqual) IsEqual(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SpecMockRecorderForIsEqual) Callback(callback func(value interface{}) (result bool)) {
	mrm.call.SetCallback(callback)
}

func NewAnnotationParserMock(ctrl *unit.Controller, options ...interface{}) *AnnotationParserMock {
	return &AnnotationParserMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *AnnotationParserMock) EXPECT() *AnnotationParserMockRecorder {
	return &AnnotationParserMockRecorder{mock: m}
}

func (m *AnnotationParserMock) Parse(source string) (annotations []interface{}) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, source)

	switch __result, __type := m.callManager.FetchCall("Parse", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].([]interface{})
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(source string) (annotations []interface{}))(source)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *AnnotationParserMockRecorder) Parse(source interface{}) *AnnotationParserMockRecorderForParse {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, source)

	return &AnnotationParserMockRecorderForParse{
		call: mr.mock.callManager.CreateCall("Parse", __params...),
	}
}

func (mrm *AnnotationParserMockRecorderForParse) Return(annotations []interface{}) {
	mrm.call.SetReturn(annotations)
}

func (mrm *AnnotationParserMockRecorderForParse) Parse(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *AnnotationParserMockRecorderForParse) Callback(callback func(source string) (annotations []interface{})) {
	mrm.call.SetCallback(callback)
}

func NewSourceParserMock(ctrl *unit.Controller, options ...interface{}) *SourceParserMock {
	return &SourceParserMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *SourceParserMock) EXPECT() *SourceParserMockRecorder {
	return &SourceParserMockRecorder{mock: m}
}

func (m *SourceParserMock) Parse(fileName string, content string) (result0 *File) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, fileName)
	__params = append(__params, content)

	switch __result, __type := m.callManager.FetchCall("Parse", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(*File)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(fileName string, content string) *File)(fileName, content)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *SourceParserMockRecorder) Parse(fileName interface{}, content interface{}) *SourceParserMockRecorderForParse {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, fileName)
	__params = append(__params, content)

	return &SourceParserMockRecorderForParse{
		call: mr.mock.callManager.CreateCall("Parse", __params...),
	}
}

func (mrm *SourceParserMockRecorderForParse) Return(result0 *File) {
	mrm.call.SetReturn(result0)
}

func (mrm *SourceParserMockRecorderForParse) Parse(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *SourceParserMockRecorderForParse) Callback(callback func(fileName string, content string) *File) {
	mrm.call.SetCallback(callback)
}

func NewMarshalerJSONMock(ctrl *unit.Controller, options ...interface{}) *MarshalerJSONMock {
	return &MarshalerJSONMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *MarshalerJSONMock) EXPECT() *MarshalerJSONMockRecorder {
	return &MarshalerJSONMockRecorder{mock: m}
}

func (m *MarshalerJSONMock) MarshalJSON() (result0 []byte, result1 error) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("MarshalJSON").Call(); __type {
	case MockCallTypeReturn:
		results := __result.([]interface{})

		if results[0] != nil {
			result0 = __result.([]interface{})[0].([]byte)
		}

		if results[1] != nil {
			result1 = __result.([]interface{})[1].(error)
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() ([]byte, error))()
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *MarshalerJSONMockRecorder) MarshalJSON() *MarshalerJSONMockRecorderForMarshalJSON {
	mr.mock.ctrl.TestingT().Helper()

	return &MarshalerJSONMockRecorderForMarshalJSON{
		call: mr.mock.callManager.CreateCall("MarshalJSON"),
	}
}

func (mrm *MarshalerJSONMockRecorderForMarshalJSON) Return(result0 []byte, result1 error) {
	mrm.call.SetReturn(result0, result1)
}

func (mrm *MarshalerJSONMockRecorderForMarshalJSON) MarshalJSON(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *MarshalerJSONMockRecorderForMarshalJSON) Callback(callback func() ([]byte, error)) {
	mrm.call.SetCallback(callback)
}
