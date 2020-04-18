package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityEqualer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityEqualer{}

	actual := NewEntityEqualer()

	ctrl.AssertEqual(expected, actual)
}

func TestEntityEqualer_Equal_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	y := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithSimpleSpecAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithSimpleSpecAndPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	y := &SimpleSpec{
		PackageName: "another",
		TypeName:    "typeName",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithSimpleSpecAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &SimpleSpec{
		TypeName: "typeName",
	}

	y := &SimpleSpec{
		TypeName: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithSimpleSpecAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	y := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: false,
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "length",
	}

	y := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "length",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithArraySpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "length",
	}

	y := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "another",
		},
		Length: "length",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithArraySpecAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "length",
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	y := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithMapSpecAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithMapSpecAndKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	y := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "another",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithMapSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	y := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "another",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Field{
		Name: "name",
		Tag:  "tag",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Field{
		Name: "name",
		Tag:  "tag",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithFieldAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Field{
		Name: "name",
		Tag:  "tag",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFieldAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Field{
		Name: "name",
	}

	y := &Field{
		Name: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFieldAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Field{
		Name: "name",
		Tag:  "tag",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Field{
		Name: "name",
		Tag:  "another",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFieldAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Field{
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Field{
		Spec: &SimpleSpec{
			TypeName: "another",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param3TypeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "result2TypeName",
				},
			},
		},
	}

	y := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param3TypeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "result2TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	y := &FuncSpec{
		IsVariadic: false,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndParamsLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param2TypeName",
				},
			},
		},
	}

	y := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndResultsLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "result2TypeName",
				},
			},
		},
	}

	y := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndParam(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "param2TypeName",
				},
			},
		},
	}

	y := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "another",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncSpecAndResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "result2TypeName",
				},
			},
		},
	}

	y := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "another",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	y := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithInterfaceSpecAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	y := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithInterfaceSpecAndFieldsLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
		},
	}

	y := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithInterfaceSpecAndFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	y := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "another",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	y := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithStructSpecAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
		},
	}

	y := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field3TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithStructSpecAndFieldsLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
		},
	}

	y := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithStructSpecAndFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
		},
	}

	y := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "another",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	y := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithImportAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Import{
		Namespace: "namespace/alias",
	}

	y := &Import{
		Alias:     "alias",
		Namespace: "namespace/alias",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithImportAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	y := &Import{
		Alias:     "another",
		Namespace: "namespace",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportAndNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	y := &Import{
		Alias:     "alias",
		Namespace: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias1",
				Namespace: "namespace1",
			},
			{
				Alias:     "alias2",
				Namespace: "namespace2",
			},
		},
	}

	y := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias1",
				Namespace: "namespace1",
			},
			{
				Alias:     "alias2",
				Namespace: "namespace2",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace1",
			},
			{
				Namespace: "namespace2",
			},
		},
	}

	y := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace2",
			},
			{
				Namespace: "namespace1",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace1",
			},
		},
	}

	y := &ImportGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace/alias",
			},
		},
	}

	y := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace/alias",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace",
			},
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	y := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "another",
				Namespace: "namespace",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithImportGroupAndNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	y := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "another",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	y := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithConstAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Const{
		Name:  "name",
		Value: "value",
	}

	y := &Const{
		Name:  "another",
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	y := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "another",
		},
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Const{
		Name:  "name",
		Value: "value",
	}

	y := &Const{
		Name:  "name",
		Value: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithConstGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{}
	y := &ConstGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithConstGroupAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name1",
				Value: "value1",
			},
			{
				Name:  "name2",
				Value: "value2",
			},
		},
	}

	y := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name2",
				Value: "value2",
			},
			{
				Name:  "name1",
				Value: "value1",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstGroupAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := &ConstGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstGroupAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithConstGroupAndConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "another",
				Value: "value",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	y := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithVarAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Var{
		Name:  "name",
		Value: "value",
	}

	y := &Var{
		Name:  "another",
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	y := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "another",
		},
		Value: "value",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Var{
		Name:  "name",
		Value: "value",
	}

	y := &Var{
		Name:  "name",
		Value: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{
		Vars: []*Var{
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
				Value: "value1",
			},
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
				Value: "value2",
			},
		},
	}

	y := &VarGroup{
		Vars: []*Var{
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
				Value: "value1",
			},
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
				Value: "value2",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithVarGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{}
	y := &VarGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithVarGroupAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Value: "value1",
			},
			{
				Name:  "name2",
				Value: "value2",
			},
		},
	}

	y := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name2",
				Value: "value2",
			},
			{
				Name:  "name1",
				Value: "value1",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithVarGroupAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := &VarGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarGroupAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithVarGroupAndVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	y := &VarGroup{
		Vars: []*Var{
			{
				Name:  "another",
				Value: "value",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithTypeAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithTypeAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Type{
		Name: "another",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithTypeAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	y := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "another",
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{
		Types: []*Type{
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
			},
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
			},
		},
	}

	y := &TypeGroup{
		Types: []*Type{
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
			},
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithTypeGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{}
	y := &TypeGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithTypeGroupAndOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{
		Types: []*Type{
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
			},
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
			},
		},
	}

	y := &TypeGroup{
		Types: []*Type{
			{
				Name: "name2",
				Spec: &SimpleSpec{
					TypeName: "typeName2",
				},
			},
			{
				Name: "name1",
				Spec: &SimpleSpec{
					TypeName: "typeName1",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithTypeGroupAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	y := &TypeGroup{}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithTypeGroupAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithTypeGroupAndType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	y := &TypeGroup{
		Types: []*Type{
			{
				Name: "another",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name:    "name",
		Content: "content",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	y := &Func{
		Name:    "name",
		Content: "content",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithFuncAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
	}

	y := &Func{
		Name: "name",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertTrue(actual)
}

func TestEntityEqualer_Equal_WithFuncAndAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
	}

	y := "y"

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
	}

	y := &Func{
		Name: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name:    "name",
		Content: "content",
	}

	y := &Func{
		Name:    "name",
		Content: "another",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndNilSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
	}

	y := &Func{
		Name: "name",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndNilRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	y := &Func{
		Name: "name",
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
	}

	y := &Func{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "another",
					},
				},
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}

func TestEntityEqualer_Equal_WithFuncAndRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	x := &Func{
		Name: "name",
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	y := &Func{
		Name: "name",
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "another",
			},
		},
	}

	actual := (&EntityEqualer{}).Equal(x, y)

	ctrl.AssertFalse(actual)
}
