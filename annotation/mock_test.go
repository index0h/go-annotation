package annotation

import (
	"github.com/index0h/go-unit/unit"
	"github.com/pkg/errors"
)

type (
	ScannerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ScannerMockRecorder struct {
		mock *ScannerMock
	}

	ScannerMockRecorderForScan struct {
		call *MockCall
	}
)

type (
	RendererMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	RendererMockRecorder struct {
		mock *RendererMock
	}

	RendererMockRecorderForRender struct {
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
	AnnotationParserMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	AnnotationParserMockRecorder struct {
		mock *AnnotationParserMock
	}

	AnnotationParserMockRecorderForSetAnnotation struct {
		call *MockCall
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
	StorageWriterMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	StorageWriterMockRecorder struct {
		mock *StorageWriterMock
	}

	StorageWriterMockRecorderForWrite struct {
		call *MockCall
	}
)

type (
	StorageCleanerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	StorageCleanerMockRecorder struct {
		mock *StorageCleanerMock
	}

	StorageCleanerMockRecorderForClean struct {
		call *MockCall
	}
)

type (
	GeneratorMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	GeneratorMockRecorder struct {
		mock *GeneratorMock
	}

	GeneratorMockRecorderForAnnotations struct {
		call *MockCall
	}

	GeneratorMockRecorderForGenerate struct {
		call *MockCall
	}
)

type (
	ImportUniquerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	ImportUniquerMockRecorder struct {
		mock *ImportUniquerMock
	}

	ImportUniquerMockRecorderForUnique struct {
		call *MockCall
	}
)

type (
	EqualerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	EqualerMockRecorder struct {
		mock *EqualerMock
	}

	EqualerMockRecorderForEqual struct {
		call *MockCall
	}
)

type (
	MarshalerMock struct {
		ctrl        *unit.Controller
		callManager *MockCallManager
	}

	MarshalerMockRecorder struct {
		mock *MarshalerMock
	}

	MarshalerMockRecorderForMarshalJSON struct {
		call *MockCall
	}
)

func NewScannerMock(ctrl *unit.Controller, options ...interface{}) *ScannerMock {
	return &ScannerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ScannerMock) EXPECT() *ScannerMockRecorder {
	return &ScannerMockRecorder{mock: m}
}

func (m *ScannerMock) Scan(storage *Storage, rootNamespace string, rootPath string, ignores ...string) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)
	__params = append(__params, rootNamespace)
	__params = append(__params, rootPath)

	for _, __param := range ignores {
		__params = append(__params, __param)
	}

	switch __result, __type := m.callManager.FetchCall("Scan", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(storage *Storage, rootNamespace string, rootPath string, ignores ...string))(storage, rootNamespace, rootPath, ignores...)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ScannerMockRecorder) Scan(storage interface{}, rootNamespace interface{}, rootPath interface{}, ignores ...interface{}) *ScannerMockRecorderForScan {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)
	__params = append(__params, rootNamespace)
	__params = append(__params, rootPath)

	for _, __param := range ignores {
		__params = append(__params, __param)
	}

	return &ScannerMockRecorderForScan{
		call: mr.mock.callManager.CreateCall("Scan", __params...),
	}
}

func (mrm *ScannerMockRecorderForScan) Return() {
	mrm.call.SetReturn()
}

func (mrm *ScannerMockRecorderForScan) Scan(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ScannerMockRecorderForScan) Callback(callback func(storage *Storage, rootNamespace string, rootPath string, ignores ...string)) {
	mrm.call.SetCallback(callback)
}

func NewRendererMock(ctrl *unit.Controller, options ...interface{}) *RendererMock {
	return &RendererMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *RendererMock) EXPECT() *RendererMockRecorder {
	return &RendererMockRecorder{mock: m}
}

func (m *RendererMock) Render(entity interface{}) (result0 string) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, entity)

	switch __result, __type := m.callManager.FetchCall("Render", __params...).Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].(string)
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(entity interface{}) string)(entity)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *RendererMockRecorder) Render(entity interface{}) *RendererMockRecorderForRender {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, entity)

	return &RendererMockRecorderForRender{
		call: mr.mock.callManager.CreateCall("Render", __params...),
	}
}

func (mrm *RendererMockRecorderForRender) Return(result0 string) {
	mrm.call.SetReturn(result0)
}

func (mrm *RendererMockRecorderForRender) Render(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *RendererMockRecorderForRender) Callback(callback func(entity interface{}) string) {
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

func (m *ValidatorMock) Validate(entity interface{}) (result0 error) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, entity)

	switch __result, __type := m.callManager.FetchCall("Validate", __params...).Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].(error)
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(entity interface{}) error)(entity)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ValidatorMockRecorder) Validate(entity interface{}) *ValidatorMockRecorderForValidate {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, entity)

	return &ValidatorMockRecorderForValidate{
		call: mr.mock.callManager.CreateCall("Validate", __params...),
	}
}

func (mrm *ValidatorMockRecorderForValidate) Return(result0 error) {
	mrm.call.SetReturn(result0)
}

func (mrm *ValidatorMockRecorderForValidate) Validate(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ValidatorMockRecorderForValidate) Callback(callback func(entity interface{}) error) {
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

func (m *AnnotationParserMock) SetAnnotation(name string, annotationType interface{}) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, name)
	__params = append(__params, annotationType)

	switch __result, __type := m.callManager.FetchCall("SetAnnotation", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(name string, annotationType interface{}))(name, annotationType)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *AnnotationParserMockRecorder) SetAnnotation(name interface{}, annotationType interface{}) *AnnotationParserMockRecorderForSetAnnotation {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, name)
	__params = append(__params, annotationType)

	return &AnnotationParserMockRecorderForSetAnnotation{
		call: mr.mock.callManager.CreateCall("SetAnnotation", __params...),
	}
}

func (mrm *AnnotationParserMockRecorderForSetAnnotation) Return() {
	mrm.call.SetReturn()
}

func (mrm *AnnotationParserMockRecorderForSetAnnotation) SetAnnotation(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *AnnotationParserMockRecorderForSetAnnotation) Callback(callback func(name string, annotationType interface{})) {
	mrm.call.SetCallback(callback)
}

func (m *AnnotationParserMock) Parse(content string) (annotations []interface{}) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, content)

	switch __result, __type := m.callManager.FetchCall("Parse", __params...).Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			annotations = __results[0].([]interface{})
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(content string) (annotations []interface{}))(content)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *AnnotationParserMockRecorder) Parse(content interface{}) *AnnotationParserMockRecorderForParse {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, content)

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

func (mrm *AnnotationParserMockRecorderForParse) Callback(callback func(content string) (annotations []interface{})) {
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
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].(*File)
		}

		return
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

func NewStorageWriterMock(ctrl *unit.Controller, options ...interface{}) *StorageWriterMock {
	return &StorageWriterMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *StorageWriterMock) EXPECT() *StorageWriterMockRecorder {
	return &StorageWriterMockRecorder{mock: m}
}

func (m *StorageWriterMock) Write(storage *Storage) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)

	switch __result, __type := m.callManager.FetchCall("Write", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(storage *Storage))(storage)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *StorageWriterMockRecorder) Write(storage interface{}) *StorageWriterMockRecorderForWrite {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)

	return &StorageWriterMockRecorderForWrite{
		call: mr.mock.callManager.CreateCall("Write", __params...),
	}
}

func (mrm *StorageWriterMockRecorderForWrite) Return() {
	mrm.call.SetReturn()
}

func (mrm *StorageWriterMockRecorderForWrite) Write(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *StorageWriterMockRecorderForWrite) Callback(callback func(storage *Storage)) {
	mrm.call.SetCallback(callback)
}

func NewStorageCleanerMock(ctrl *unit.Controller, options ...interface{}) *StorageCleanerMock {
	return &StorageCleanerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *StorageCleanerMock) EXPECT() *StorageCleanerMockRecorder {
	return &StorageCleanerMockRecorder{mock: m}
}

func (m *StorageCleanerMock) Clean(storage *Storage) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)

	switch __result, __type := m.callManager.FetchCall("Clean", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(storage *Storage))(storage)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *StorageCleanerMockRecorder) Clean(storage interface{}) *StorageCleanerMockRecorderForClean {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, storage)

	return &StorageCleanerMockRecorderForClean{
		call: mr.mock.callManager.CreateCall("Clean", __params...),
	}
}

func (mrm *StorageCleanerMockRecorderForClean) Return() {
	mrm.call.SetReturn()
}

func (mrm *StorageCleanerMockRecorderForClean) Clean(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *StorageCleanerMockRecorderForClean) Callback(callback func(storage *Storage)) {
	mrm.call.SetCallback(callback)
}

func NewGeneratorMock(ctrl *unit.Controller, options ...interface{}) *GeneratorMock {
	return &GeneratorMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *GeneratorMock) EXPECT() *GeneratorMockRecorder {
	return &GeneratorMockRecorder{mock: m}
}

func (m *GeneratorMock) Annotations() (result0 map[string]interface{}) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("Annotations").Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].(map[string]interface{})
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func() map[string]interface{})()
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *GeneratorMockRecorder) Annotations() *GeneratorMockRecorderForAnnotations {
	mr.mock.ctrl.TestingT().Helper()

	return &GeneratorMockRecorderForAnnotations{
		call: mr.mock.callManager.CreateCall("Annotations"),
	}
}

func (mrm *GeneratorMockRecorderForAnnotations) Return(result0 map[string]interface{}) {
	mrm.call.SetReturn(result0)
}

func (mrm *GeneratorMockRecorderForAnnotations) Annotations(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *GeneratorMockRecorderForAnnotations) Callback(callback func() map[string]interface{}) {
	mrm.call.SetCallback(callback)
}

func (m *GeneratorMock) Generate(param0 *Application) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	switch __result, __type := m.callManager.FetchCall("Generate", __params...).Call(); __type {
	case MockCallTypeReturn:
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		__result.(func(*Application))(param0)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *GeneratorMockRecorder) Generate(param0 interface{}) *GeneratorMockRecorderForGenerate {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, param0)

	return &GeneratorMockRecorderForGenerate{
		call: mr.mock.callManager.CreateCall("Generate", __params...),
	}
}

func (mrm *GeneratorMockRecorderForGenerate) Return() {
	mrm.call.SetReturn()
}

func (mrm *GeneratorMockRecorderForGenerate) Generate(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *GeneratorMockRecorderForGenerate) Callback(callback func(*Application)) {
	mrm.call.SetCallback(callback)
}

func NewImportUniquerMock(ctrl *unit.Controller, options ...interface{}) *ImportUniquerMock {
	return &ImportUniquerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *ImportUniquerMock) EXPECT() *ImportUniquerMockRecorder {
	return &ImportUniquerMockRecorder{mock: m}
}

func (m *ImportUniquerMock) Unique(list []*Import) (result0 []*Import) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, list)

	switch __result, __type := m.callManager.FetchCall("Unique", __params...).Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].([]*Import)
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(list []*Import) []*Import)(list)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *ImportUniquerMockRecorder) Unique(list interface{}) *ImportUniquerMockRecorderForUnique {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, list)

	return &ImportUniquerMockRecorderForUnique{
		call: mr.mock.callManager.CreateCall("Unique", __params...),
	}
}

func (mrm *ImportUniquerMockRecorderForUnique) Return(result0 []*Import) {
	mrm.call.SetReturn(result0)
}

func (mrm *ImportUniquerMockRecorderForUnique) Unique(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *ImportUniquerMockRecorderForUnique) Callback(callback func(list []*Import) []*Import) {
	mrm.call.SetCallback(callback)
}

func NewEqualerMock(ctrl *unit.Controller, options ...interface{}) *EqualerMock {
	return &EqualerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *EqualerMock) EXPECT() *EqualerMockRecorder {
	return &EqualerMockRecorder{mock: m}
}

func (m *EqualerMock) Equal(x interface{}, y interface{}) (result0 bool) {
	m.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, x)
	__params = append(__params, y)

	switch __result, __type := m.callManager.FetchCall("Equal", __params...).Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].(bool)
		}

		return
	case MockCallTypePanic:
		panic(__result)
	case MockCallTypeCallback:
		return __result.(func(x interface{}, y interface{}) bool)(x, y)
	default:
		panic(errors.New("Unknown mock call type, you should regenerate mock"))
	}
}

func (mr *EqualerMockRecorder) Equal(x interface{}, y interface{}) *EqualerMockRecorderForEqual {
	mr.mock.ctrl.TestingT().Helper()

	__params := []interface{}{}
	__params = append(__params, x)
	__params = append(__params, y)

	return &EqualerMockRecorderForEqual{
		call: mr.mock.callManager.CreateCall("Equal", __params...),
	}
}

func (mrm *EqualerMockRecorderForEqual) Return(result0 bool) {
	mrm.call.SetReturn(result0)
}

func (mrm *EqualerMockRecorderForEqual) Equal(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *EqualerMockRecorderForEqual) Callback(callback func(x interface{}, y interface{}) bool) {
	mrm.call.SetCallback(callback)
}

func NewMarshalerMock(ctrl *unit.Controller, options ...interface{}) *MarshalerMock {
	return &MarshalerMock{
		ctrl:        ctrl,
		callManager: NewMockCallManager(ctrl, options...),
	}
}

func (m *MarshalerMock) EXPECT() *MarshalerMockRecorder {
	return &MarshalerMockRecorder{mock: m}
}

func (m *MarshalerMock) MarshalJSON() (result0 []byte, result1 error) {
	m.ctrl.TestingT().Helper()

	switch __result, __type := m.callManager.FetchCall("MarshalJSON").Call(); __type {
	case MockCallTypeReturn:
		__results := __result.([]interface{})

		if __results[0] != nil {
			result0 = __results[0].([]byte)
		}

		if __results[1] != nil {
			result1 = __results[1].(error)
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

func (mr *MarshalerMockRecorder) MarshalJSON() *MarshalerMockRecorderForMarshalJSON {
	mr.mock.ctrl.TestingT().Helper()

	return &MarshalerMockRecorderForMarshalJSON{
		call: mr.mock.callManager.CreateCall("MarshalJSON"),
	}
}

func (mrm *MarshalerMockRecorderForMarshalJSON) Return(result0 []byte, result1 error) {
	mrm.call.SetReturn(result0, result1)
}

func (mrm *MarshalerMockRecorderForMarshalJSON) MarshalJSON(value interface{}) {
	mrm.call.SetPanic(value)
}

func (mrm *MarshalerMockRecorderForMarshalJSON) Callback(callback func() ([]byte, error)) {
	mrm.call.SetCallback(callback)
}
