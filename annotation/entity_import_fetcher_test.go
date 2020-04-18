package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewImportFetcher(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importUniquer := NewImportUniquerMock(ctrl)

	actual := NewEntityImportFetcher(importUniquer)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(importUniquer, actual.importUniquer)
}

func TestNewImportFetcher_WithNilImportUniquer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewEntityImportFetcher, nil).
		ExpectPanic(NewErrorMessageConstraint("Variable 'importUniquer' must be not nil"))
}

func TestImportFetcher_Fetch_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &SimpleSpec{
		TypeName: "typeName",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithSimpleSpecAndPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithSimpleSpecAndPackageNameAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &SimpleSpec{
		PackageName: "another",
		TypeName:    "typeName",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(file.ImportGroups[0].Imports).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithArraySpecAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "packageName.MyConst + 1",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(file.ImportGroups[0].Imports).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithArraySpecAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithMapSpecAndKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "keyPackageName",
						Namespace: "keyNamespace",
					},
					{
						Alias:     "valuePackageName",
						Namespace: "valueNamespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &MapSpec{
		Key: &SimpleSpec{
			PackageName: "keyPackageName",
			TypeName:    "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(
			[]*Import{
				file.ImportGroups[0].Imports[0],
			},
		).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithMapSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "keyPackageName",
						Namespace: "keyNamespace",
					},
					{
						Alias:     "valuePackageName",
						Namespace: "valueNamespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[1],
	}

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			PackageName: "valuePackageName",
			TypeName:    "valueTypeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(
			[]*Import{
				file.ImportGroups[0].Imports[1],
			},
		).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithMapSpecAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "keyPackageName",
						Namespace: "keyNamespace",
					},
					{
						Alias:     "valuePackageName",
						Namespace: "valueNamespace",
					},
				},
			},
		},
	}

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Field{
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithFieldAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	entity := &Field{
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "paramPackageName",
						Namespace: "paramNamespace",
					},
					{
						Alias:     "resultPackageName",
						Namespace: "resultNamespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
		file.ImportGroups[0].Imports[1],
	}

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "paramPackageName",
					TypeName:    "paramTypeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "resultPackageName",
					TypeName:    "resultTypeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithFuncSpecAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "paramPackageName",
						Namespace: "paramNamespace",
					},
					{
						Alias:     "resultPackageName",
						Namespace: "resultNamespace",
					},
				},
			},
		},
	}

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "paramTypeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "resultTypeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithInterfaceSpecAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithStructSpecAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
		Value: "value",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithConstAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "packageName.MyConst + 1",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithConstAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name: "name",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
				Value: "value",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithConstGroupAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
				Value: "packageName.MyConst + 1",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithConstGroupAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
				Value: "value",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithVarAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "packageName.MyVar + 1",
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithVarAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithVarGroupAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
				Value: "packageName.MyVar + 1",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithVarGroupAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithTypeAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
	}

	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithTypeGroupAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "namespace/packageName",
					},
				},
			},
		},
	}

	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "paramPackageName",
						Namespace: "paramNamespace",
					},
					{
						Alias:     "resultPackageName",
						Namespace: "resultNamespace",
					},
					{
						Alias:     "relatedPackageName",
						Namespace: "relatedNamespace",
					},
				},
			},
		},
	}

	expected := []*Import{
		file.ImportGroups[0].Imports[0],
		file.ImportGroups[0].Imports[1],
		file.ImportGroups[0].Imports[2],
	}

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: "paramPackageName",
						TypeName:    "paramTypeName",
					},
				},
			},
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						PackageName: "resultPackageName",
						TypeName:    "resultTypeName",
					},
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				PackageName: "relatedPackageName",
				TypeName:    "relatedTypeName",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique(
			[]*Import{
				file.ImportGroups[0].Imports[0],
				file.ImportGroups[0].Imports[1],
			},
		).
		Return(
			[]*Import{
				file.ImportGroups[0].Imports[0],
				file.ImportGroups[0].Imports[1],
			},
		)

	importUniquer.
		EXPECT().
		Unique(expected).
		Return(expected)

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_Fetch_WithFuncAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "paramPackageName",
						Namespace: "paramNamespace",
					},
					{
						Alias:     "resultPackageName",
						Namespace: "resultNamespace",
					},
					{
						Alias:     "relatedPackageName",
						Namespace: "relatedNamespace",
					},
				},
			},
		},
	}

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "resultTypeName",
					},
				},
			},
		},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	importUniquer.
		EXPECT().
		Unique([]*Import{}).
		Return([]*Import{})

	actual := entityImportFetcher.Fetch(file, entity)

	ctrl.AssertEmpty(actual)
}

func TestImportFetcher_Fetch_WithUnknownType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	entity := "invalid"

	ctrl.Subtest("").
		Call((&EntityImportFetcher{importUniquer: NewEntityImportUniquer()}).Fetch, file, entity).
		ExpectPanic(NewErrorMessageConstraint("Can't fetch entity with type: 'string'"))
}

func TestImportFetcher_fetchFromContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "fmt",
					},
					{
						Namespace: "strconv",
					},
				},
			},
		},
	}

	expected := []*Import{
		{
			Namespace: "fmt",
		},
		{
			Namespace: "strconv",
		},
	}

	content := "fmt.Println(\"Hello world\" + strconv.Itoa(5))"

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.fetchFromContent(file, content)

	ctrl.AssertEqual(expected, actual)
}

func TestImportFetcher_fetchFromContent_WithEmptyImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
	}

	content := "// comment"

	importUniquer := NewImportUniquerMock(ctrl)

	entityImportFetcher := &EntityImportFetcher{importUniquer: importUniquer}

	actual := entityImportFetcher.fetchFromContent(file, content)

	ctrl.AssertEmpty(actual)
}
