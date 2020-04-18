package annotation

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityCloner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityCloner{}

	actual := NewEntityCloner()

	ctrl.AssertEqual(expected, actual)
}

func TestEntityCloner_Clone_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
		Length: "length",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Value, actual.(*ArraySpec).Value)
}

func TestEntityCloner_Clone_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Key, actual.(*MapSpec).Key)
	ctrl.AssertNotSame(entity.Value, actual.(*MapSpec).Value)
}

func TestEntityCloner_Clone_WithField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Name:    "fieldName",
		Tag:     "fieldTag",
		Comment: "fieldComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "fieldAnnotation",
			},
		},
		Spec: &SimpleSpec{
			TypeName: "fieldSpecTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*Field).Annotations[0])
	ctrl.AssertNotSame(entity.Spec, actual.(*Field).Spec)
}

func TestEntityCloner_Clone_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "funcParamSpecTypeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "funcResultSpecTypeName",
				},
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Params[0].Spec, actual.(*FuncSpec).Params[0].Spec)
	ctrl.AssertNotSame(entity.Params[0].Spec.(*ArraySpec).Value, actual.(*FuncSpec).Params[0].Spec.(*ArraySpec).Value)
	ctrl.AssertNotSame(entity.Results[0].Spec, actual.(*FuncSpec).Results[0].Spec)
}

func TestEntityCloner_Clone_WithFuncSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
}

func TestEntityCloner_Clone_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldSpecTypeName",
				},
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Fields[0].Spec, actual.(*InterfaceSpec).Fields[0].Spec)
}

func TestEntityCloner_Clone_WithInterfaceSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
}

func TestEntityCloner_Clone_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldSpecTypeName",
				},
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Fields[0].Spec, actual.(*StructSpec).Fields[0].Spec)
}

func TestEntityCloner_Clone_WithStructSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
}

func TestEntityCloner_Clone_WithImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importAnnotation",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*Import).Annotations[0])
}

func TestEntityCloner_Clone_WithImportAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Namespace: "namespace",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{
		Comment: "importGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importGroupAnnotation",
			},
		},
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

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*ImportGroup).Annotations[0])
	ctrl.AssertNotSame(entity.Imports[0], actual.(*ImportGroup).Imports[0])
	ctrl.AssertNotSame(entity.Imports[1], actual.(*ImportGroup).Imports[1])
}

func TestEntityCloner_Clone_WithImportGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name:    "constName",
		Value:   "iota",
		Comment: "const\ncomment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constAnnotation",
			},
		},
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*Const).Annotations[0])
}

func TestEntityCloner_Clone_WithConstAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name:  "constName",
		Value: "constValue",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{
		Comment: "constGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constGroupAnnotation",
			},
		},
		Consts: []*Const{
			{
				Name:  "const1Name",
				Value: "const1Value",
			},
			{
				Name:  "const2Name",
				Value: "const2Value",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*ConstGroup).Annotations[0])
	ctrl.AssertNotSame(entity.Consts[0], actual.(*ConstGroup).Consts[0])
	ctrl.AssertNotSame(entity.Consts[1], actual.(*ConstGroup).Consts[1])
}

func TestEntityCloner_Clone_WithConstGroupAndFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name:    "name",
		Comment: "comment",
		Value:   "value",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varAnnotation",
			},
		},
		Spec: &SimpleSpec{
			TypeName: "varTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Spec, actual.(*Var).Spec)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*Var).Annotations[0])
}

func TestEntityCloner_Clone_WithVarAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "varTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Spec, actual.(*Var).Spec)
}

func TestEntityCloner_Clone_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name:  "var1Name",
				Value: "var1Value",
			},
			{
				Name:  "var2Name",
				Value: "var2Value",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*VarGroup).Annotations[0])
	ctrl.AssertNotSame(entity.Vars[0], actual.(*VarGroup).Vars[0])
	ctrl.AssertNotSame(entity.Vars[1], actual.(*VarGroup).Vars[1])
}

func TestEntityCloner_Clone_WithVarGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name:    "name",
		Comment: "comment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeAnnotation",
			},
		},
		Spec: &SimpleSpec{
			TypeName: "typeTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Spec, actual.(*Type).Spec)
}

func TestEntityCloner_Clone_WithTypeAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeTypeName",
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Spec, actual.(*Type).Spec)
}

func TestEntityCloner_Clone_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "type1Name",
				Spec: &SimpleSpec{
					TypeName: "type1TypeName",
				},
			},
			{
				Name: "type2Name",
				Spec: &SimpleSpec{
					TypeName: "type2TypeName",
				},
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*TypeGroup).Annotations[0])
	ctrl.AssertNotSame(entity.Types[0], actual.(*TypeGroup).Types[0])
	ctrl.AssertNotSame(entity.Types[0].Spec, actual.(*TypeGroup).Types[0].Spec)
	ctrl.AssertNotSame(entity.Types[1], actual.(*TypeGroup).Types[1])
	ctrl.AssertNotSame(entity.Types[1].Spec, actual.(*TypeGroup).Types[1].Spec)
}

func TestEntityCloner_Clone_WithTypeGroupAndFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
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
					Spec: &SimpleSpec{
						TypeName: "specParamSpecTypeName",
					},
				},
			},
		},
		Related: &Field{
			Name:    "relatedName",
			Comment: "relatedComment",
			Spec: &SimpleSpec{
				TypeName: "relatedSpecTypeName",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Spec.Params[0].Spec, actual.(*Func).Spec.Params[0].Spec)
	ctrl.AssertNotSame(entity.Related.Spec, actual.(*Func).Related.Spec)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*Func).Annotations[0])
}

func TestEntityCloner_Clone_WithFuncAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Comment:     "file\nComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "fileAnnotation",
			},
		},
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup1\nComment",
			},
			{
				Comment: "importGroup2\nComment",
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup1\nComment",
			},
			{
				Comment: "constGroup2\nComment",
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup1\nComment",
			},
			{
				Comment: "varGroup2\nComment",
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup1\nComment",
			},
			{
				Comment: "typeGroup2\nComment",
			},
		},
		Funcs: []*Func{
			{
				Name: "funcName1",
			},
			{
				Name: "funcName2",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Annotations[0], actual.(*File).Annotations[0])
	ctrl.AssertNotSame(entity.ImportGroups[0], actual.(*File).ImportGroups[0])
	ctrl.AssertNotSame(entity.ImportGroups[1], actual.(*File).ImportGroups[1])
	ctrl.AssertNotSame(entity.ConstGroups[0], actual.(*File).ConstGroups[0])
	ctrl.AssertNotSame(entity.ConstGroups[1], actual.(*File).ConstGroups[1])
	ctrl.AssertNotSame(entity.VarGroups[0], actual.(*File).VarGroups[0])
	ctrl.AssertNotSame(entity.VarGroups[1], actual.(*File).VarGroups[1])
	ctrl.AssertNotSame(entity.TypeGroups[0], actual.(*File).TypeGroups[0])
	ctrl.AssertNotSame(entity.TypeGroups[1], actual.(*File).TypeGroups[1])
	ctrl.AssertNotSame(entity.Funcs[0], actual.(*File).Funcs[0])
	ctrl.AssertNotSame(entity.Funcs[1], actual.(*File).Funcs[1])
}

func TestEntityCloner_Clone_WithFileAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name:      "namespace/alias",
		Path:      "/namespace/path",
		IsIgnored: true,
		Files: []*File{
			{
				Name:        "file1Name.go",
				PackageName: "file1PackageName",
			},
			{
				Name:        "file2Name.go",
				PackageName: "file2PackageName",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Files[0], actual.(*Namespace).Files[0])
	ctrl.AssertNotSame(entity.Files[1], actual.(*Namespace).Files[1])
}

func TestEntityCloner_Clone_WithNamespaceAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertEqual(entity, actual)
	ctrl.AssertNotSame(entity, actual)
}

func TestEntityCloner_Clone_WithStorage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace1/alias",
				Path: "/namespace1/path",
			},
			{
				Name: "namespace2/alias",
				Path: "/namespace2/path",
			},
		},
	}

	actual := (&EntityCloner{}).Clone(entity)

	ctrl.AssertNotSame(entity, actual)
	ctrl.AssertNotSame(entity.Namespaces[0], actual.(*Storage).Namespaces[0])
	ctrl.AssertNotSame(entity.Namespaces[1], actual.(*Storage).Namespaces[1])
}

func TestEntityCloner_Clone_WithUnknownType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := "data"

	ctrl.Subtest("").
		Call((&EntityCloner{}).Clone, entity).
		ExpectPanic(
			NewErrorMessageConstraint("Can't clone entity with type: 'string'"),
		)
}

func TestEntityCloner_cloneAnnotations_WithMarshalJSONPanic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	err := errors.New("TestError")

	annotation := NewMarshalerMock(ctrl)

	annotations := []interface{}{
		annotation,
	}

	annotation.
		EXPECT().
		MarshalJSON().
		Return(nil, err)

	ctrl.Subtest("").
		Call((&EntityCloner{}).cloneAnnotations, annotations).
		ExpectPanic(ctrl.Type(&json.MarshalerError{}))
}

func TestEntityCloner_cloneAnnotations_WithUnmarshalJSONPanic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	marshaled := []byte(`"data"`)

	annotation := NewMarshalerMock(ctrl)

	annotations := []interface{}{
		annotation,
	}

	annotation.
		EXPECT().
		MarshalJSON().
		Return(marshaled, nil)

	ctrl.Subtest("").
		Call((&EntityCloner{}).cloneAnnotations, annotations).
		ExpectPanic(ctrl.Type(&json.UnmarshalTypeError{}))
}
