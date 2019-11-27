package annotation

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestFunc_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
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

	model.Validate()
}

func TestFunc_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "funcName",
	}

	model.Validate()
}

func TestFunc_Validate_WithEmptyFuncName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFunc_Validate_WithInvalidFuncName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "+invalid",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFunc_Validate_WithInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			IsVariadic: true,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params' must be not empty for variadic %T", model.Spec))
}

func TestFunc_Validate_WithInvalidRelatedField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Name: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFunc_Validate_WithInvalidRelatedFieldSpecType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Spec: &FuncSpec{},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Related.Spec.(%T)' has invalid type for %T", model.Spec, model),
		)
}

func TestFunc_Validate_WithInvalidRelatedFieldSpecPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
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
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Related.Spec.(%T).PackageName' must be empty for %T",
				model.Related.Spec,
				model,
			),
		)
}

func TestFunc_Validate_WithInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name:    "funcName",
		Spec:    &FuncSpec{},
		Content: "[[",
	}

	ctrl.Subtest("").
		Call(model.Validate).
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

	model := &Func{
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

	actual := model.String()

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

	model := &Func{
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

	actual := model.String()

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

	model := &Func{
		Name:    name,
		Content: content,
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
	}

	actual := model.String()

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

	model := &Func{
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

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertSame(clonedFuncSpecParamSpec, actual.(*Func).Spec.Params[0].Spec)
	ctrl.AssertSame(clonedFuncSpecResultSpec, actual.(*Func).Spec.Results[0].Spec)
	ctrl.AssertSame(clonedRelatedSpec, actual.(*Func).Related.Spec)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*Func).Annotations[0])
}

func TestFunc_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Func{
		Name: "funcName",
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
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

	model := &Func{
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

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestFunc_FetchImports_WithoutImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	relatedSpec := NewSpecMock(ctrl)

	model := &Func{
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

	actual := model.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestFunc_FetchImports_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	model := &Func{
		Name: "funcName",
	}

	actual := model.FetchImports(file)

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

	model := &Func{
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

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertSame(expectedContent, model.Content)
}

func TestFunc_RenameImports_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Func{
		Name: "funcName",
	}

	model.RenameImports(oldAlias, newAlias)
}

func TestFunc_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestFunc_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	model := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
