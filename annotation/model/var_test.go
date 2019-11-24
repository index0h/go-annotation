package model

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestVar_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &StructSpec{},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestVar_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestVar_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "+invalid",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestVar_Validate_WithNilSpecAndEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("*model.Var must have not nil 'Spec' or not empty 'Value'"))
}

func TestVar_Validate_WithInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name:  "name",
		Value: "[invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestVar_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestVar_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
			Length: -100,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
}

func TestVar_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestVar_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Var{
		Name: "name",
		Spec: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: *model.SpecMock"))
}

func TestVar_String_WithCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	comment := "type\ncomment"
	value := "value"
	specSpecString := "specSpecString"
	expected := `// type
// comment
var name specSpecString = value
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name:    name,
		Comment: comment,
		Value:   value,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestVar_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	comment := "type\ncomment"
	specSpecString := "specSpecString"
	expected := `// type
// comment
var name specSpecString
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name:    name,
		Comment: comment,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestVar_String_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	value := "value"
	specSpecString := "specSpecString"
	expected := `var name specSpecString = value
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name:  name,
		Value: value,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestVar_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	specSpecString := "specSpecString"
	expected := `var name specSpecString
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: name,
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestVar_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name:    "name",
		Comment: "comment",
		Value:   "value",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		Clone().
		Return(NewSpecMock(ctrl))

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Spec, actual.(*Var).Spec)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*Var).Annotations[0])
}

func TestVar_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		Clone().
		Return(NewSpecMock(ctrl))

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Spec, actual.(*Var).Spec)
}

func TestVar_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "packageName",
			Namespace: "namespace",
		},
	}

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := modelValue.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestVar_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestVar_RenameImports_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"
	expectedValue := "(newPackageName.value + 5) * iota"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name:  "name",
		Value: "(oldPackageName.value + 5) * iota",
		Spec:  specSpec,
	}

	specSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertSame(expectedValue, modelValue.Value)
}

func TestVar_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestVar_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Var{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
