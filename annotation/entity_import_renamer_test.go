package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityImportRenamer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityImportRenamer{}

	actual := NewEntityImportRenamer()

	ctrl.AssertEqual(expected, actual)
}

func TestEntityImportRenamer_Rename_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &SimpleSpec{
		PackageName: oldAlias,
		TypeName:    "typeName",
	}

	expected := &SimpleSpec{
		PackageName: newAlias,
		TypeName:    "typeName",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithSimpleSpecAndNotRenamed(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &SimpleSpec{
		TypeName: "typeName",
	}

	expected := &SimpleSpec{
		TypeName: "typeName",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithArraySpecAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "oldPackageName.VarName + 1",
	}

	expected := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "newPackageName.VarName + 1",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &MapSpec{
		Key: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "keyTypeName",
		},
		Value: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "valueTypeName",
		},
	}

	expected := &MapSpec{
		Key: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "keyTypeName",
		},
		Value: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "valueTypeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Field{
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &Field{
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFuncSpecAndEmptyFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &FuncSpec{}

	expected := &FuncSpec{}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithInterfaceSpecAndFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "Func",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: &SimpleSpec{
								PackageName: oldAlias,
								TypeName:    "typeName",
							},
						},
					},
				},
			},
		},
	}

	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "Func",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: &SimpleSpec{
								PackageName: newAlias,
								TypeName:    "typeName",
							},
						},
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Import{
		Alias:     oldAlias,
		Namespace: "namespace",
	}

	expected := &Import{
		Alias:     newAlias,
		Namespace: "namespace",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithImportAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Import{
		Namespace: "namespace/oldPackageName",
	}

	expected := &Import{
		Alias:     newAlias,
		Namespace: "namespace/oldPackageName",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     oldAlias,
				Namespace: "namespace",
			},
		},
	}

	expected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     newAlias,
				Namespace: "namespace",
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithImportGroupAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace/oldPackageName",
			},
		},
	}

	expected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     newAlias,
				Namespace: "namespace/oldPackageName",
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Const{
		Name:  "constName",
		Value: "(oldPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &Const{
		Name:  "constName",
		Value: "(newPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "(oldPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "(newPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Var{
		Name:  "varName",
		Value: "(oldPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &Var{
		Name:  "varName",
		Value: "(newPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name:  "varName",
				Value: "(oldPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &VarGroup{
		Vars: []*Var{
			{
				Name:  "varName",
				Value: "(newPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Type{
		Name: "typeName",
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &Type{
		Name: "typeName",
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	expected := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Func{
		Name:    "funcName",
		Content: "return (oldPackageName.data + 1) * 5",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: oldAlias,
						TypeName:    "paramTypeName",
					},
				},
			},
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: oldAlias,
						TypeName:    "resultTypeName",
					},
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				PackageName: oldAlias,
				TypeName:    "relatedTypeName",
			},
		},
	}

	expected := &Func{
		Name:    "funcName",
		Content: "return (newPackageName.data + 1) * 5",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: newAlias,
						TypeName:    "paramTypeName",
					},
				},
			},
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: newAlias,
						TypeName:    "resultTypeName",
					},
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				PackageName: newAlias,
				TypeName:    "relatedTypeName",
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFunAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &Func{
		Name: "funcName",
	}

	expected := &Func{
		Name: "funcName",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName",
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName",
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content: `package packageName

import namespace/oldPackageName

type typeName oldPackageName.TypeName`,
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content: `package packageName

import namespace/oldPackageName

type typeName newPackageName.TypeName`,
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     oldAlias,
						Namespace: "namespace",
					},
				},
			},
		},
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     newAlias,
						Namespace: "namespace",
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name: "constName",
						Spec: &SimpleSpec{
							PackageName: oldAlias,
							TypeName:    "typeName",
						},
						Value: "value",
					},
				},
			},
		},
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name: "constName",
						Spec: &SimpleSpec{
							PackageName: newAlias,
							TypeName:    "typeName",
						},
						Value: "value",
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name: "varName",
						Spec: &SimpleSpec{
							PackageName: oldAlias,
							TypeName:    "typeName",
						},
						Value: "value",
					},
				},
			},
		},
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name: "varName",
						Spec: &SimpleSpec{
							PackageName: newAlias,
							TypeName:    "typeName",
						},
						Value: "value",
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: &SimpleSpec{
							PackageName: oldAlias,
							TypeName:    "typeName",
						},
					},
				},
			},
		},
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: &SimpleSpec{
							PackageName: newAlias,
							TypeName:    "typeName",
						},
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithFileAndFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			{
				Name:        "funcName",
				Annotations: nil,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: &SimpleSpec{
								PackageName: newAlias,
								TypeName:    "paramTypeName",
							},
						},
					},
				},
			},
		},
	}

	expected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			{
				Name:        "funcName",
				Annotations: nil,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: &SimpleSpec{
								PackageName: newAlias,
								TypeName:    "paramTypeName",
							},
						},
					},
				},
			},
		},
	}

	(&EntityImportRenamer{}).Rename(entity, oldAlias, newAlias)

	ctrl.AssertEqual(expected, entity)
}

func TestEntityImportRenamer_Rename_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	ctrl.Subtest("").
		Call((&EntityImportRenamer{}).Rename, &FuncSpec{}, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestEntityImportRenamer_Rename_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	ctrl.Subtest("").
		Call((&EntityImportRenamer{}).Rename, &FuncSpec{}, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestEntityImportRenamer_Rename_WithInvalidType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	ctrl.Subtest("").
		Call((&EntityImportRenamer{}).Rename, "invalid", oldAlias, newAlias).
		ExpectPanic(NewErrorMessageConstraint("Can't rename entity with type: 'string'"))
}
