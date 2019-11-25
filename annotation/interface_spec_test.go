package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestInterfaceSpec_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	model.Validate()
}

func TestInterfaceSpec_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	model.Validate()
}

func TestInterfaceSpec_Validate_WithoutMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{}

	model.Validate()
}

func TestInterfaceSpec_Validate_WithNilField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestInterfaceSpec_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "+invalid",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestInterfaceSpec_Validate_WithSimpleSpecValueAndMethodName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0].Name' must be empty for 'Fields[0].Spec' type *SimpleSpec",
			),
		)
}

func TestInterfaceSpec_Validate_WithSimpleSpecValueAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName:  "typeName",
					IsPointer: true,
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Fields[0].Spec.(%T).IsPointer' must be 'false'", model.Fields[0].Spec),
		)
}

func TestInterfaceSpec_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &FuncSpec{
					Params: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestInterfaceSpec_Validate_WithFuncSpecValueAndWithoutMethodName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0].Name' must be not empty for 'Fields[0].Spec' type *FuncSpec",
			),
		)
}

func TestInterfaceSpec_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' has invalid type %T", model.Fields[0].Spec))
}

func TestInterfaceSpec_Validate_WithInvalidField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "+invalid",
				Spec: &FuncSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestInterfaceSpec_String_WithCommentAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	name := "name"
	methodSpecString := "(methodSpecString)"
	expected := `interface{
// comment
// here
name(methodSpecString)
}`

	methodSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name:    name,
				Comment: comment,
				Spec:    methodSpec,
			},
		},
	}

	methodSpec.
		EXPECT().
		String().
		Return(methodSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestInterfaceSpec_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	methodSpecString := "methodSpecString"
	expected := `interface{
// comment
// here
methodSpecString
}`

	methodSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Comment: comment,
				Spec:    methodSpec,
			},
		},
	}

	methodSpec.
		EXPECT().
		String().
		Return(methodSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestInterfaceSpec_String_WithName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	methodSpecString := "(methodSpecString)"
	expected := `interface{
name(methodSpecString)
}`

	methodSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: name,
				Spec: methodSpec,
			},
		},
	}

	methodSpec.
		EXPECT().
		String().
		Return(methodSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestInterfaceSpec_String_WithoutMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "interface{}"

	model := &InterfaceSpec{}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestInterfaceSpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	methodSpec := NewSpecMock(ctrl)
	clonedMethodSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "methodAnnotation",
					},
				},
				Spec: methodSpec,
			},
		},
	}

	methodSpec.
		EXPECT().
		Clone().
		Return(clonedMethodSpec)

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertSame(clonedMethodSpec, actual.(*InterfaceSpec).Fields[0].Spec)
	ctrl.AssertNotSame(model.Fields[0].Annotations[0], actual.(*InterfaceSpec).Fields[0].Annotations[0])
}

func TestInterfaceSpec_Clone_WithoutMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &InterfaceSpec{}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
}

func TestInterfaceSpec_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "packageName",
			Namespace: "namespace",
		},
	}

	valueSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: valueSpec,
			},
		},
	}

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestInterfaceSpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: valueSpec,
			},
		},
	}

	valueSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestInterfaceSpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &InterfaceSpec{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestInterfaceSpec_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	model := &InterfaceSpec{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
