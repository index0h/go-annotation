package model

import (
	"github.com/index0h/go-unit/unit"
	"github.com/pkg/errors"
)

type TestAnnotation struct {
	Name string
}

type (
	StringerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	StringerMockRecorder struct {
		mock *StringerMock
	}

	StringerMockRecorderForString struct {
		call *MockCall
	}
)

type (
	ValidatorMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ValidatorMockRecorder struct {
		mock *ValidatorMock
	}

	ValidatorMockRecorderForValidate struct {
		call *MockCall
	}
)

type (
	ClonerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ClonerMockRecorder struct {
		mock *ClonerMock
	}

	ClonerMockRecorderForClone struct {
		call *MockCall
	}
)

type (
	ImportsFetcherMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ImportsFetcherMockRecorder struct {
		mock *ImportsFetcherMock
	}

	ImportsFetcherMockRecorderForFetchImports struct {
		call *MockCall
	}
)

type (
	ImportsRenamerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ImportsRenamerMockRecorder struct {
		mock *ImportsRenamerMock
	}

	ImportsRenamerMockRecorderForRenameImports struct {
		call *MockCall
	}
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

type (
	UnmarshalerJSONMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	UnmarshalerJSONMockRecorder struct {
		mock *UnmarshalerJSONMock
	}

	UnmarshalerJSONMockRecorderForUnmarshalJSON struct {
		call *MockCall
	}
)

type (
	MarshalerAndUnmarshalerJSONMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	MarshalerAndUnmarshalerJSONMockRecorder struct {
		mock *MarshalerAndUnmarshalerJSONMock
	}

	MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON struct {
		call *MockCall
	}

	MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON struct {
		call *MockCall
	}
)

func NewStringerMock(ctrl *unit.Controller, options ...interface{}) *StringerMock {
	return &StringerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *StringerMock) EXPECT() *StringerMockRecorder {
	return &StringerMockRecorder{mock: m}
}

func (m *StringerMock) String() (result string) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("String").Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(string)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() (result string))()
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *StringerMockRecorder) String() *StringerMockRecorderForString {
	mr.mock.ctrl.TestingT().Helper()

	return &StringerMockRecorderForString{
		call: mr.mock.callManager.CreateCall("String"),
	}
}

func (mrm *StringerMockRecorderForString) Return(result string) {
	mrm.call.SetReturn(result)
}

func (mrm *StringerMockRecorderForString) String(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *StringerMockRecorderForString) Callback(callback func() (result string)) {
	mrm.call.SetCallback(callback)
}

func NewValidatorMock(ctrl *unit.Controller, options ...interface{}) *ValidatorMock {
	return &ValidatorMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ValidatorMock) EXPECT() *ValidatorMockRecorder {
	return &ValidatorMockRecorder{mock: m}
}

func (m *ValidatorMock) Validate() {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("Validate").Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func())()
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ValidatorMockRecorder) Validate() *ValidatorMockRecorderForValidate {
	mr.mock.ctrl.TestingT().Helper()

	return &ValidatorMockRecorderForValidate{
		call: mr.mock.callManager.CreateCall("Validate"),
	}
}

func (mrm *ValidatorMockRecorderForValidate) Return() {
	mrm.call.SetReturn()
}

func (mrm *ValidatorMockRecorderForValidate) Validate(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ValidatorMockRecorderForValidate) Callback(callback func()) {
	mrm.call.SetCallback(callback)
}

func NewClonerMock(ctrl *unit.Controller, options ...interface{}) *ClonerMock {
	return &ClonerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ClonerMock) EXPECT() *ClonerMockRecorder {
	return &ClonerMockRecorder{mock: m}
}

func (m *ClonerMock) Clone() (result interface{}) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("Clone").Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(interface{})
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() (result interface{}))()
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ClonerMockRecorder) Clone() *ClonerMockRecorderForClone {
	mr.mock.ctrl.TestingT().Helper()

	return &ClonerMockRecorderForClone{
		call: mr.mock.callManager.CreateCall("Clone"),
	}
}

func (mrm *ClonerMockRecorderForClone) Return(result interface{}) {
	mrm.call.SetReturn(result)
}

func (mrm *ClonerMockRecorderForClone) Clone(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ClonerMockRecorderForClone) Callback(callback func() (result interface{})) {
	mrm.call.SetCallback(callback)
}

func NewImportsFetcherMock(ctrl *unit.Controller, options ...interface{}) *ImportsFetcherMock {
	return &ImportsFetcherMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ImportsFetcherMock) EXPECT() *ImportsFetcherMockRecorder {
	return &ImportsFetcherMockRecorder{mock: m}
}

func (m *ImportsFetcherMock) FetchImports(file *File) (result0 []*Import) {
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ImportsFetcherMockRecorder) FetchImports(file interface{}) *ImportsFetcherMockRecorderForFetchImports {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, file)

	return &ImportsFetcherMockRecorderForFetchImports{
		call: mr.mock.callManager.CreateCall("FetchImports", __params...),
	}
}

func (mrm *ImportsFetcherMockRecorderForFetchImports) Return(result0 []*Import) {
	mrm.call.SetReturn(result0)
}

func (mrm *ImportsFetcherMockRecorderForFetchImports) FetchImports(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ImportsFetcherMockRecorderForFetchImports) Callback(callback func(file *File) []*Import) {
	mrm.call.SetCallback(callback)
}

func NewImportsRenamerMock(ctrl *unit.Controller, options ...interface{}) *ImportsRenamerMock {
	return &ImportsRenamerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ImportsRenamerMock) EXPECT() *ImportsRenamerMockRecorder {
	return &ImportsRenamerMockRecorder{mock: m}
}

func (m *ImportsRenamerMock) RenameImports(oldAlias string, newAlias string) {
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ImportsRenamerMockRecorder) RenameImports(oldAlias interface{}, newAlias interface{}) *ImportsRenamerMockRecorderForRenameImports {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, oldAlias)
	__params = append(__params, newAlias)

	return &ImportsRenamerMockRecorderForRenameImports{
		call: mr.mock.callManager.CreateCall("RenameImports", __params...),
	}
}

func (mrm *ImportsRenamerMockRecorderForRenameImports) Return() {
	mrm.call.SetReturn()
}

func (mrm *ImportsRenamerMockRecorderForRenameImports) RenameImports(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ImportsRenamerMockRecorderForRenameImports) Callback(callback func(oldAlias string, newAlias string)) {
	mrm.call.SetCallback(callback)
}

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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
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

func NewUnmarshalerJSONMock(ctrl *unit.Controller, options ...interface{}) *UnmarshalerJSONMock {
	return &UnmarshalerJSONMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *UnmarshalerJSONMock) EXPECT() *UnmarshalerJSONMockRecorder {
	return &UnmarshalerJSONMockRecorder{mock: m}
}

func (m *UnmarshalerJSONMock) UnmarshalJSON(param0 []byte) (result0 error) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	switch __result, __type := m.callManager.FetchCall("UnmarshalJSON", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(error)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func([]byte) error)(param0)
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *UnmarshalerJSONMockRecorder) UnmarshalJSON(param0 interface{}) *UnmarshalerJSONMockRecorderForUnmarshalJSON {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	return &UnmarshalerJSONMockRecorderForUnmarshalJSON{
		call: mr.mock.callManager.CreateCall("UnmarshalJSON", __params...),
	}
}

func (mrm *UnmarshalerJSONMockRecorderForUnmarshalJSON) Return(result0 error) {
	mrm.call.SetReturn(result0)
}

func (mrm *UnmarshalerJSONMockRecorderForUnmarshalJSON) UnmarshalJSON(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *UnmarshalerJSONMockRecorderForUnmarshalJSON) Callback(callback func([]byte) error) {
	mrm.call.SetCallback(callback)
}

func NewMarshalerAndUnmarshalerJSONMock(ctrl *unit.Controller, options ...interface{}) *MarshalerAndUnmarshalerJSONMock {
	return &MarshalerAndUnmarshalerJSONMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *MarshalerAndUnmarshalerJSONMock) EXPECT() *MarshalerAndUnmarshalerJSONMockRecorder {
	return &MarshalerAndUnmarshalerJSONMockRecorder{mock: m}
}

func (m *MarshalerAndUnmarshalerJSONMock) MarshalJSON() (result0 []byte, result1 error) {
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

		return result0, result1
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() ([]byte, error))()
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *MarshalerAndUnmarshalerJSONMockRecorder) MarshalJSON() *MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON {
	mr.mock.ctrl.TestingT().Helper()

	return &MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON{
		call: mr.mock.callManager.CreateCall("MarshalJSON"),
	}
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON) Return(result0 []byte, result1 error) {
	mrm.call.SetReturn(result0, result1)
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON) MarshalJSON(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForMarshalJSON) Callback(callback func() ([]byte, error)) {
	mrm.call.SetCallback(callback)
}

func (m *MarshalerAndUnmarshalerJSONMock) UnmarshalJSON(param0 []byte) (result0 error) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	switch __result, __type := m.callManager.FetchCall("UnmarshalJSON", __params...).Call(); __type {
	case MockCallTypeReturn:
		return __result.([]interface{})[0].(error)
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func([]byte) error)(param0)
	default:
		panic(errors.Errorf("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *MarshalerAndUnmarshalerJSONMockRecorder) UnmarshalJSON(param0 interface{}) *MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	return &MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON{
		call: mr.mock.callManager.CreateCall("UnmarshalJSON", __params...),
	}
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON) Return(result0 error) {
	mrm.call.SetReturn(result0)
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON) UnmarshalJSON(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *MarshalerAndUnmarshalerJSONMockRecorderForUnmarshalJSON) Callback(callback func([]byte) error) {
	mrm.call.SetCallback(callback)
}
