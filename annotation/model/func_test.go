package model

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestFunc_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name:    "funcName",
		Content: "funcContent",
		Comment: "funcComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	modelValue.Validate()
}

func TestFunc_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestFunc_Validate_WithEmptyFuncName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFunc_Validate_WithInvalidFuncName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "+invalid",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFunc_Validate_WithNilSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' must be not nil"))
}

func TestFunc_Validate_WithInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			IsVariadic: true,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params' must be not empty for variadic *model.FuncSpec"))
}

func TestFunc_Validate_WithInvalidRelatedField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Name: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFunc_Validate_WithInvalidRelatedFieldSpecType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Spec: &FuncSpec{},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Related.Spec.(*model.FuncSpec)' has invalid type for *model.Func",
			),
		)
}

func TestFunc_Validate_WithInvalidRelatedFieldSpecPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Spec: &SimpleSpec{
				PackageName: "packageName",
				TypeName:    "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Related.Spec.(*model.SimpleSpec).PackageName' must be empty for *model.Func",
			),
		)
}

func TestFunc_Validate_WithInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name:    "funcName",
		Spec:    &FuncSpec{},
		Content: "[[",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestFunc_String_WithCommentAndRelatedNameAndRelatedComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "funcName"
	content := "funcContent"
	comment := "func\ncomment"
	relatedName := "relatedName"
	relatedComment := "related\ncomment"
	funcSpecParamString := "funcSpecParamString"
	funcSpecResultString := "funcSpecResultString"
	relatedSpecString := "relatedSpecString"

	expected := `// func
// comment
func (
// related
// comment
relatedName relatedSpecString) funcName(param funcSpecParamString) (result funcSpecResultString) {
funcContent
}
`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name:    name,
		Content: content,
		Comment: comment,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "param",
					Spec: funcSpecParamSpec,
				},
			},
			Results: []*Field{
				{
					Name: "result",
					Spec: funcSpecResultSpec,
				},
			},
		},
		Related: &Field{
			Name:    relatedName,
			Comment: relatedComment,
			Spec:    relatedSpec,
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamString)

	funcSpecResultSpec.
		EXPECT().
		String().
		Return(funcSpecResultString)

	relatedSpec.
		EXPECT().
		String().
		Return(relatedSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFunc_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "funcName"
	content := "funcContent"
	comment := "func\ncomment"
	funcSpecParamString := "funcSpecParamString"
	funcSpecResultString := "funcSpecResultString"

	expected := `// func
// comment
func funcName(param funcSpecParamString) (result funcSpecResultString) {
funcContent
}
`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResultSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name:    name,
		Content: content,
		Comment: comment,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "param",
					Spec: funcSpecParamSpec,
				},
			},
			Results: []*Field{
				{
					Name: "result",
					Spec: funcSpecResultSpec,
				},
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamString)

	funcSpecResultSpec.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFunc_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "funcName"
	content := "funcContent"

	expected := `func funcName() {
funcContent
}
`

	modelValue := &Func{
		Name:    name,
		Content: content,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFunc_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	clonedFuncSpecParamSpec := NewSpecMock(ctrl)
	clonedFuncSpecResultSpec := NewSpecMock(ctrl)
	clonedRelatedSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name:    "funcName",
		Comment: "funcComment",
		Content: "funcContent",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "param",
					Spec: funcSpecParamSpec,
				},
			},
			Results: []*Field{
				{
					Name: "result",
					Spec: funcSpecResultSpec,
				},
			},
		},
		Related: &Field{
			Name:    "relatedName",
			Comment: "relatedComment",
			Spec:    relatedSpec,
		},
	}

	funcSpecParamSpec.
		EXPECT().
		Clone().
		Return(clonedFuncSpecParamSpec)

	funcSpecResultSpec.
		EXPECT().
		Clone().
		Return(clonedFuncSpecResultSpec)

	relatedSpec.
		EXPECT().
		Clone().
		Return(clonedRelatedSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedFuncSpecParamSpec, actual.(*Func).Spec.Params[0].Spec)
	ctrl.AssertSame(clonedFuncSpecResultSpec, actual.(*Func).Spec.Results[0].Spec)
	ctrl.AssertSame(clonedRelatedSpec, actual.(*Func).Related.Spec)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*Func).Annotations[0])
}

func TestFunc_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Spec, actual.(*Func).Spec)
}

func TestFunc_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "funcParamPackageName",
			Namespace: "funcParamNamespace",
		},
		{
			Alias:     "funcResultPackageName",
			Namespace: "funcResultNamespace",
		},
		{
			Alias:     "relatedPackageName",
			Namespace: "relatedNamespace",
		},
	}

	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: funcParamSpec,
				},
			},
			Results: []*Field{
				{
					Spec: funcResultSpec,
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: relatedSpec,
		},
	}

	funcParamSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[0]})

	funcResultSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[1]})

	relatedSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[2]})

	actual := modelValue.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestFunc_FetchImports_WithoutImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: funcParamSpec,
				},
			},
			Results: []*Field{
				{
					Spec: funcResultSpec,
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: relatedSpec,
		},
	}

	funcParamSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	funcResultSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	relatedSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestFunc_FetchImports_WithoutRelatedAndWithEmptyFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestFunc_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"
	expectedContent := "return (newPackageName.data + 1) * 5"

	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	modelValue := &Func{
		Name:    "funcName",
		Content: "return (oldPackageName.data + 1) * 5",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: funcParamSpec,
				},
			},
			Results: []*Field{
				{
					Spec: funcResultSpec,
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: relatedSpec,
		},
	}

	funcParamSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcResultSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	relatedSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertSame(expectedContent, modelValue.Content)
}

func TestFunc_RenameImports_WithoutRelatedAndWithEmptySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestFunc_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestFunc_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	modelValue := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
