package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityContainsChecker(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	equaler := NewEqualerMock(ctrl)

	expected := &EntityContainsChecker{
		equaler: equaler,
	}

	actual := NewEntityContainsChecker(equaler)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(expected.equaler, actual.equaler)
}

func TestEntityContainsChecker_Contains_WithImportAndImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias2",
		Namespace: "namespace2",
	}

	collection := &ImportGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndImportGroupAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Namespace: "namespace2/alias2",
	}

	collection := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias1",
				Namespace: "namespace1",
			},
			{
				Alias:     "alias2",
				Namespace: "namespace2/alias2",
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndImportGroupAndEmptyImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias2",
		Namespace: "namespace2",
	}

	collection := &ImportGroup{}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndImportGroupAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias3",
		Namespace: "namespace3",
	}

	collection := &ImportGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &ConstGroup{
		Consts: []*Const{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndConstGroupAndEmptyConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &ConstGroup{}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndConstGroupAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
		Value: "value3",
	}

	collection := &ConstGroup{
		Consts: []*Const{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &VarGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndVarGroupAndEmptyVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &VarGroup{}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndVarGroupAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
		Value: "value3",
	}

	collection := &VarGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
	}

	collection := &TypeGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndTypeGroupAndEmptyTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
	}

	collection := &TypeGroup{}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndTypeGroupAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
	}

	collection := &TypeGroup{
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

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias2",
		Namespace: "namespace2",
	}

	collection := &File{
		ImportGroups: []*ImportGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndFileAndRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Namespace: "namespace2/alias2",
	}

	collection := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "alias1",
						Namespace: "namespace1",
					},
					{
						Alias:     "alias2",
						Namespace: "namespace2/alias2",
					},
				},
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndFileAndEmptyImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias2",
		Namespace: "namespace2",
	}

	collection := &File{
		ImportGroups: []*ImportGroup{},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithImportAndFileAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias3",
		Namespace: "namespace3",
	}

	collection := &File{
		ImportGroups: []*ImportGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &File{
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndFileAndEmptyConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &File{
		ConstGroups: []*ConstGroup{},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithConstAndFileAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
		Value: "value3",
	}

	collection := &File{
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &File{
		VarGroups: []*VarGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndFileAndEmptyVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value2",
	}

	collection := &File{
		VarGroups: []*VarGroup{},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithVarAndFileAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
		Value: "value3",
	}

	collection := &File{
		VarGroups: []*VarGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
	}

	collection := &File{
		TypeGroups: []*TypeGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndFileAndEmptyTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
	}

	collection := &File{
		TypeGroups: []*TypeGroup{},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithTypeAndFileAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name3",
		Spec: &SimpleSpec{
			TypeName: "typeName3",
		},
	}

	collection := &File{
		TypeGroups: []*TypeGroup{
			{
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
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithFuncAndFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "name1",
					Spec: &SimpleSpec{
						TypeName: "typeName1",
					},
				},
			},
		},
	}

	collection := &File{
		Funcs: []*Func{
			{
				Name: "funcName",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "name1",
							Spec: &SimpleSpec{
								TypeName: "typeName1",
							},
						},
					},
				},
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertTrue(actual)
}

func TestEntityContainsChecker_Contains_WithFuncAndFileAndEmptyFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "name1",
					Spec: &SimpleSpec{
						TypeName: "typeName1",
					},
				},
			},
		},
	}

	collection := &File{}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithFuncAndFileAndNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "name1",
					Spec: &SimpleSpec{
						TypeName: "typeName1",
					},
				},
			},
		},
	}

	collection := &File{
		Funcs: []*Func{
			{
				Name: "anotherFuncName",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "name1",
							Spec: &SimpleSpec{
								TypeName: "typeName1",
							},
						},
					},
				},
			},
		},
	}

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}

func TestEntityContainsChecker_Contains_WithInvalidCollectionType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	collection := "invalid"

	EntityContainsChecker := &EntityContainsChecker{
		equaler: NewEntityEqualer(),
	}

	actual := EntityContainsChecker.Contains(collection, entity)

	ctrl.AssertFalse(actual)
}
