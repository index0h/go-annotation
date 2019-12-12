package annotation

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestVar_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	model.Validate()
}

func TestVar_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	model.Validate()
}

func TestVar_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
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

	model.Validate()
}

func TestVar_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &StructSpec{},
	}

	model.Validate()
}

func TestVar_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &InterfaceSpec{},
	}

	model.Validate()
}

func TestVar_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &FuncSpec{},
	}

	model.Validate()
}

func TestVar_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestVar_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "+invalid",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestVar_Validate_WithNilSpecAndEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("%T must have not nil 'Spec' or not empty 'Value'", model))
}

func TestVar_Validate_WithInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name:  "name",
		Value: "[invalid",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestVar_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestVar_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &ArraySpec{},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestVar_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestVar_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestVar_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Var{
		Name: "name",
		Spec: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: %T", model.Spec))
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

	model := &Var{
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

	actual := model.String()

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

	model := &Var{
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

	actual := model.String()

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

	model := &Var{
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

	actual := model.String()

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

	model := &Var{
		Name: name,
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVar_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)

	model := &Var{
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

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Spec, actual.(*Var).Spec)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*Var).Annotations[0])
}

func TestVar_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)

	model := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		Clone().
		Return(NewSpecMock(ctrl))

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Spec, actual.(*Var).Spec)
}

func TestVar_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name:  "name",
		Spec:  model1Spec,
		Value: "value",
	}

	model2 := &Var{
		Name:  "name",
		Spec:  model2Spec,
		Value: "value",
	}

	model1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVar_EqualSpec_WithoutSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Var{
		Name:  "name",
		Value: "value",
	}

	model2 := &Var{
		Name:  "name",
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVar_EqualSpec_WithoutValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name: "name",
		Spec: model1Spec,
	}

	model2 := &Var{
		Name: "name",
		Spec: model2Spec,
	}

	model1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVar_EqualSpec_WithAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name: "name",
		Spec: model1Spec,
	}

	model2 := "model2"

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVar_EqualSpec_WithName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name:  "name1",
		Spec:  model1Spec,
		Value: "value",
	}

	model2 := &Var{
		Name:  "name2",
		Spec:  model2Spec,
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVar_EqualSpec_WithSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name:  "name",
		Spec:  model1Spec,
		Value: "value",
	}

	model2 := &Var{
		Name:  "name",
		Spec:  model2Spec,
		Value: "value",
	}

	model1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec)).
		Return(false)

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVar_EqualSpec_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &Var{
		Name:  "name",
		Spec:  model1Spec,
		Value: "value1",
	}

	model2 := &Var{
		Name:  "name",
		Spec:  model2Spec,
		Value: "value2",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
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

	model := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestVar_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	model := &Var{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestVar_RenameImports_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"
	expectedValue := "(newPackageName.value + 5) * iota"

	specSpec := NewSpecMock(ctrl)

	model := &Var{
		Name:  "name",
		Value: "(oldPackageName.value + 5) * iota",
		Spec:  specSpec,
	}

	specSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertSame(expectedValue, model.Value)
}

func TestVar_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	model := &Var{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
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

	model := &Var{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
