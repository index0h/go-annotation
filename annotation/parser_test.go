package annotation

import (
	"encoding/json"
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewParser).
		ExpectResult(
			&Parser{
				annotations: map[string]interface{}{
					"FileIsGenerated": FileIsGeneratedAnnotation(false),
				},
			},
		)
}

func TestParser_AddAnnotation(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call((&Parser{annotations: map[string]interface{}{}}).AddAnnotation, "Annotation", 5).
		ExpectResult(&Parser{annotations: map[string]interface{}{"Annotation": 5}})

	ctrl.Subtest("WithEmptyNameAndNegativeResult").
		Call((&Parser{annotations: map[string]interface{}{}}).AddAnnotation, "", 5).
		ExpectPanic(NewNotEmptyError("name"))

	ctrl.Subtest("WithDuplicateNameAndNegativeResult").
		Call((&Parser{annotations: map[string]interface{}{"Annotation": 5}}).AddAnnotation, "Annotation", 5).
		ExpectPanic(NewErrorf("Annotation '%s' already registered", "Annotation"))
}

func TestParser_Process_WithPackage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type PackageAnnotation struct {
		Data string
	}

	content := `// PackageComment
// @Package({"Data": "Package"})
package packageName
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
						Comment: `PackageComment
@Package({"Data": "Package"})`,
						PackageName: "packageName",
						Annotations: []interface{}{PackageAnnotation{Data: "Package"}},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Package", PackageAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithOneImportAndSameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
import "import_namespace"
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:       "import_namespace",
										Namespace:   "import_namespace",
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithOneImportAndCustomAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
import alias "import_namespace"
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:       "alias",
										Namespace:   "import_namespace",
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithOneImportAndAutoAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
import "import_namespace/alias"
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:       "alias",
										Namespace:   "import_namespace/alias",
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithEmptyImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
import (
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports:     []*Import{},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithImportGroupAndSameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	type ImportAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
import (
	// ImportComment
	// @Import({"Data": "Import"})
	"import_namespace"
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:     "import_namespace",
										Namespace: "import_namespace",
										Comment: `ImportComment
@Import({"Data": "Import"})`,
										Annotations: []interface{}{ImportAnnotation{Data: "Import"}},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		AddAnnotation("Import", ImportAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithImportGroupAndCustomAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @Import({"Data": "ImportGroup"})
// ImportComment2
import (
	// ImportComment
	// @Import({"Data": "Import"})
	alias "import_namespace"
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@Import({"Data": "ImportGroup"})
ImportComment2`,
								Annotations: []interface{}{ImportAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "import_namespace",
										Comment: `ImportComment
@Import({"Data": "Import"})`,
										Annotations: []interface{}{ImportAnnotation{Data: "Import"}},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Import", ImportAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithImportWithImportGroupAndAutoAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ImportGroupAnnotation struct {
		Data string
	}

	type ImportAnnotation struct {
		Data string
	}

	content := `package packageName
// ImportGroupComment
// @ImportGroup({"Data": "ImportGroup"})
// ImportComment2
import (
	// ImportComment
	// @Import({"Data": "Import"})
	alias "import_namespace/alias"
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{Name: "file.go", Content: content},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ImportGroups: []*ImportGroup{
							{
								Comment: `ImportGroupComment
@ImportGroup({"Data": "ImportGroup"})
ImportComment2`,
								Annotations: []interface{}{ImportGroupAnnotation{Data: "ImportGroup"}},
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "import_namespace/alias",
										Comment: `ImportComment
@Import({"Data": "Import"})`,
										Annotations: []interface{}{ImportAnnotation{Data: "Import"}},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ImportGroup", ImportGroupAnnotation{}).
		AddAnnotation("Import", ImportAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneConstAndCustomType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const Const myType = 5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name:        "Const",
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneConstAndIntTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const Const = 5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name:        "Const",
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneConstAndFloatTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const Const = 5.5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name:        "Const",
										Spec:        &SimpleSpec{TypeName: "float64", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneConstAndStringTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const Const = "data"
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name:        "Const",
										Spec:        &SimpleSpec{TypeName: "string", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithEmptyConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts:      []*Const{},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithConstGroupAndCustomType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	type ConstAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
	// ConstComment
	// @Const({"Data": "Const"})
	Const myType = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name: "Const",
										Comment: `ConstComment
@Const({"Data": "Const"})`,
										Annotations: []interface{}{ConstAnnotation{Data: "Const"}},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		AddAnnotation("Const", ConstAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithConstConstAndIntTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	type ConstAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
	// ConstComment
	// @Const({"Data": "Const"})
	Const = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name: "Const",
										Comment: `ConstComment
@Const({"Data": "Const"})`,
										Annotations: []interface{}{ConstAnnotation{Data: "Const"}},
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		AddAnnotation("Const", ConstAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithConstGroupAndFloatTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	type ConstAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
	// ConstComment
	// @Const({"Data": "Const"})
	Const = 5.5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name: "Const",
										Comment: `ConstComment
@Const({"Data": "Const"})`,
										Annotations: []interface{}{ConstAnnotation{Data: "Const"}},
										Spec:        &SimpleSpec{TypeName: "float64", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		AddAnnotation("Const", ConstAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithConstGroupAndStringTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	type ConstAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
	// ConstComment
	// @Const({"Data": "Const"})
	Const = "data"
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name: "Const",
										Comment: `ConstComment
@Const({"Data": "Const"})`,
										Annotations: []interface{}{ConstAnnotation{Data: "Const"}},
										Spec:        &SimpleSpec{TypeName: "string", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		AddAnnotation("Const", ConstAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithConstGroupAndStringTypeByPreviousValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type ConstGroupAnnotation struct {
		Data string
	}

	type ConstAnnotation struct {
		Data string
	}

	content := `package packageName
// ConstGroupComment
// @ConstGroup({"Data": "ConstGroup"})
const (
	// ConstComment
	// @Const({"Data": "Const"})
	Const myType = "data"
	CopyConst
	AnotherConst = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						ConstGroups: []*ConstGroup{
							{
								Comment: `ConstGroupComment
@ConstGroup({"Data": "ConstGroup"})`,
								Annotations: []interface{}{ConstGroupAnnotation{Data: "ConstGroup"}},
								Consts: []*Const{
									{
										Name: "Const",
										Comment: `ConstComment
@Const({"Data": "Const"})`,
										Annotations: []interface{}{ConstAnnotation{Data: "Const"}},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
									{
										Name:        "CopyConst",
										Annotations: []interface{}{},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
									{
										Name:        "AnotherConst",
										Annotations: []interface{}{},
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("ConstGroup", ConstGroupAnnotation{}).
		AddAnnotation("Const", ConstAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneVarAndCustomType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var Var myType = 5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name:        "Var",
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneVarAndIntTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var Var = 5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name:        "Var",
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneVarAndFloatTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var Var = 5.5
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name:        "Var",
										Spec:        &SimpleSpec{TypeName: "float64", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneVarAndStringTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var Var = "data"
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name:        "Var",
										Spec:        &SimpleSpec{TypeName: "string", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneVarAndStringTypeByPreviousValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var Var, CopyVar myType = NewType(), 5+1
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name:        "Var",
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
										Annotations: []interface{}{},
									},
									{
										Name:        "CopyVar",
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithEmptyVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars:        []*Var{},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithVarGroupAndCustomType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	type VarAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
	// VarComment
	// @Var({"Data": "Var"})
	Var myType = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name: "Var",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		AddAnnotation("Var", VarAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithVarVarAndIntTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	type VarAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
	// VarComment
	// @Var({"Data": "Var"})
	Var = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name: "Var",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		AddAnnotation("Var", VarAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithVarGroupAndFloatTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	type VarAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
	// VarComment
	// @Var({"Data": "Var"})
	Var = 5.5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name: "Var",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "float64", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		AddAnnotation("Var", VarAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithVarGroupAndStringTypeByValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	type VarAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
	// VarComment
	// @Var({"Data": "Var"})
	Var = "data"
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name: "Var",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "string", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		AddAnnotation("Var", VarAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithVarGroupAndStringTypeByPreviousValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type VarGroupAnnotation struct {
		Data string
	}

	type VarAnnotation struct {
		Data string
	}

	content := `package packageName
// VarGroupComment
// @VarGroup({"Data": "VarGroup"})
var (
	// VarComment
	// @Var({"Data": "Var"})
	Var, CopyVar myType
	AnotherVar = 5
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						VarGroups: []*VarGroup{
							{
								Comment: `VarGroupComment
@VarGroup({"Data": "VarGroup"})`,
								Annotations: []interface{}{VarGroupAnnotation{Data: "VarGroup"}},
								Vars: []*Var{
									{
										Name: "Var",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
									{
										Name: "CopyVar",
										Comment: `VarComment
@Var({"Data": "Var"})`,
										Annotations: []interface{}{VarAnnotation{Data: "Var"}},
										Spec:        &SimpleSpec{TypeName: "myType", IsPointer: false},
									},
									{
										Name:        "AnotherVar",
										Annotations: []interface{}{},
										Spec:        &SimpleSpec{TypeName: "int", IsPointer: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("VarGroup", VarGroupAnnotation{}).
		AddAnnotation("Var", VarAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithPointerToSelectorType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type *my_package.myType
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &SimpleSpec{
											PackageName: "my_package",
											TypeName:    "myType",
											IsPointer:   true,
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithArrayType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type [5]myType
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &ArraySpec{
											Value: &SimpleSpec{
												TypeName:  "myType",
												IsPointer: false,
											},
											Length:        5,
											IsFixedLength: true,
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithMapType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type map[myType]int
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &MapSpec{
											Key:   &SimpleSpec{TypeName: "myType"},
											Value: &SimpleSpec{TypeName: "int"},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndWithoutArgumentsAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func ()
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params:  []*Field{},
											Results: []*Field{},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndArgumentsAndResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func (first firstType, second secondType) (third thirdType, fourth fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "firstType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "second",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{
												{
													Name:        "third",
													Spec:        &SimpleSpec{TypeName: "thirdType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "fourth",
													Spec:        &SimpleSpec{TypeName: "fourthType"},
													Annotations: []interface{}{},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndWithoutArgumentsAndResultsWithoutNames(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func () (thirdType, fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params: []*Field{},
											Results: []*Field{
												{
													Spec:        &SimpleSpec{TypeName: "thirdType"},
													Annotations: []interface{}{},
												},
												{
													Spec:        &SimpleSpec{TypeName: "fourthType"},
													Annotations: []interface{}{},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndArgumentsSameTypeAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func (first, second secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "second",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndArgumentsEllipsisAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func (first firstType, seconds ...secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "firstType"},
													Annotations: []interface{}{},
												},
												{
													Name: "seconds",
													Spec: &ArraySpec{
														Value:      &SimpleSpec{TypeName: "secondType"},
														IsEllipsis: true,
													},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithFuncTypeAndFieldsAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	type ThirdAnnotation struct {
		Data string
	}

	type FourthAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type func (
	// FirstComment
	// @First({"Data": "First"})
	first firstType,

	// SecondComment
	// @Second({"Data": "Second"})
	second secondType,
) (
	// ThirdComment
	// @Third({"Data": "Third"})
	third thirdType,

	// FourthComment
	// @Fourth({"Data": "Fourth"})
	fourth fourthType,
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name: "first",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec:        &SimpleSpec{TypeName: "firstType"},
												},
												{
													Name: "second",
													Comment: `SecondComment
@Second({"Data": "Second"})`,
													Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
													Spec:        &SimpleSpec{TypeName: "secondType"},
												},
											},
											Results: []*Field{
												{
													Name: "third",
													Comment: `ThirdComment
@Third({"Data": "Third"})`,
													Annotations: []interface{}{ThirdAnnotation{Data: "Third"}},
													Spec:        &SimpleSpec{TypeName: "thirdType"},
												},
												{
													Name: "fourth",
													Comment: `FourthComment
@Fourth({"Data": "Fourth"})`,
													Annotations: []interface{}{FourthAnnotation{Data: "Fourth"}},
													Spec:        &SimpleSpec{TypeName: "fourthType"},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		AddAnnotation("Third", ThirdAnnotation{}).
		AddAnnotation("Fourth", FourthAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithStructWithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type struct{}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &StructSpec{
											Fields: []*Field{},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithStructWithFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type struct{
	first  *myType "firstTag"
	second int     "secondTag"
}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name:        "first",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name:        "second",
													Tag:         "secondTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName: "int",
													},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithStructWithFieldTypeFromPrevious(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type struct{
	first, second *myType "firstTag"
}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name:        "first",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name:        "second",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithStructWithFieldsAndAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type struct{
	// FirstComment
	// @First({"Data": "First"})
	first  *myType "firstTag"
	// SecondComment
	// @Second({"Data": "Second"})
	second int "secondTag"
}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name: "first",
													Tag:  "firstTag",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name: "second",
													Tag:  "secondTag",
													Comment: `SecondComment
@Second({"Data": "Second"})`,
													Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
													Spec: &SimpleSpec{
														TypeName: "int",
													},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithInterfaceWithoutMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type interface{}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &InterfaceSpec{
											Methods: []*Field{},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithInterfaceWithMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type interface{
	First(first int) error
	Second(second int) error
}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &InterfaceSpec{
											Methods: []*Field{
												{
													Name:        "First",
													Annotations: []interface{}{},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name:        "first",
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
												{
													Name:        "Second",
													Annotations: []interface{}{},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name:        "second",
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithOneTypeWithInterfaceWithMethodsAndComments(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type FirstArgumentAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type Type interface{
	// FirstComment
	// @First({"Data": "First"})
	First(
		// FirstArgumentComment
		// @FirstArgument({"Data": "FirstArgument"})
		first int,
	) error
}
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Spec: &InterfaceSpec{
											Methods: []*Field{
												{
													Name: "First",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name: "first",
																Comment: `FirstArgumentComment
@FirstArgument({"Data": "FirstArgument"})`,
																Annotations: []interface{}{
																	FirstArgumentAnnotation{Data: "FirstArgument"},
																},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
											},
										},
										Annotations: []interface{}{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("FirstArgument", FirstArgumentAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithPointerToSelectorType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type *my_package.myType
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type", Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &SimpleSpec{
											PackageName: "my_package",
											TypeName:    "myType",
											IsPointer:   true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithArrayType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type [5]myType
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type", Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &ArraySpec{
											Value: &SimpleSpec{
												TypeName:  "myType",
												IsPointer: false,
											},
											Length:        5,
											IsFixedLength: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithMapType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type map[myType]int
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &MapSpec{
											Key:   &SimpleSpec{TypeName: "myType"},
											Value: &SimpleSpec{TypeName: "int"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndWithoutArgumentsAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func ()
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params:  []*Field{},
											Results: []*Field{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndArgumentsAndResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func (first firstType, second secondType) (third thirdType, fourth fourthType)
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "firstType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "second",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{
												{
													Name:        "third",
													Spec:        &SimpleSpec{TypeName: "thirdType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "fourth",
													Spec:        &SimpleSpec{TypeName: "fourthType"},
													Annotations: []interface{}{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndWithoutArgumentsAndResultsWithoutNames(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func () (thirdType, fourthType)
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params: []*Field{},
											Results: []*Field{
												{
													Spec:        &SimpleSpec{TypeName: "thirdType"},
													Annotations: []interface{}{},
												},
												{
													Spec:        &SimpleSpec{TypeName: "fourthType"},
													Annotations: []interface{}{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndArgumentsSameTypeAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func (first, second secondType)
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
												{
													Name:        "second",
													Spec:        &SimpleSpec{TypeName: "secondType"},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndArgumentsEllipsisAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func (first firstType, seconds ...secondType)
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name:        "first",
													Spec:        &SimpleSpec{TypeName: "firstType"},
													Annotations: []interface{}{},
												},
												{
													Name: "seconds",
													Spec: &ArraySpec{
														Value:      &SimpleSpec{TypeName: "secondType"},
														IsEllipsis: true,
													},
													Annotations: []interface{}{},
												},
											},
											Results: []*Field{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithFuncTypeAndFieldsAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	type ThirdAnnotation struct {
		Data string
	}

	type FourthAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type func (
		// FirstComment
		// @First({"Data": "First"})
		first firstType,
	
		// SecondComment
		// @Second({"Data": "Second"})
		second secondType,
	) (
		// ThirdComment
		// @Third({"Data": "Third"})
		third thirdType,
	
		// FourthComment
		// @Fourth({"Data": "Fourth"})
		fourth fourthType,
	)
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &FuncSpec{
											Params: []*Field{
												{
													Name: "first",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec:        &SimpleSpec{TypeName: "firstType"},
												},
												{
													Name: "second",
													Comment: `SecondComment
@Second({"Data": "Second"})`,
													Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
													Spec:        &SimpleSpec{TypeName: "secondType"},
												},
											},
											Results: []*Field{
												{
													Name: "third",
													Comment: `ThirdComment
@Third({"Data": "Third"})`,
													Annotations: []interface{}{ThirdAnnotation{Data: "Third"}},
													Spec:        &SimpleSpec{TypeName: "thirdType"},
												},
												{
													Name: "fourth",
													Comment: `FourthComment
@Fourth({"Data": "Fourth"})`,
													Annotations: []interface{}{FourthAnnotation{Data: "Fourth"}},
													Spec:        &SimpleSpec{TypeName: "fourthType"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		AddAnnotation("Third", ThirdAnnotation{}).
		AddAnnotation("Fourth", FourthAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithStructWithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type struct{}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &StructSpec{
											Fields: []*Field{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithStructWithFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type struct{
		first  *myType "firstTag"
		second int     "secondTag"
	}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name:        "first",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name:        "second",
													Tag:         "secondTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName: "int",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithStructWithFieldTypeFromPrevious(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type struct{
		first, second *myType "firstTag"
	}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name:        "first",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name:        "second",
													Tag:         "firstTag",
													Annotations: []interface{}{},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithStructWithFieldsAndComments(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type struct{
		// FirstComment
		// @First({"Data": "First"})
		first  *myType "firstTag"
		// SecondComment
		// @Second({"Data": "Second"})
		second int "secondTag"
	}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &StructSpec{
											Fields: []*Field{
												{
													Name: "first",
													Tag:  "firstTag",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec: &SimpleSpec{
														TypeName:  "myType",
														IsPointer: true,
													},
												},
												{
													Name: "second",
													Tag:  "secondTag",
													Comment: `SecondComment
@Second({"Data": "Second"})`,
													Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
													Spec: &SimpleSpec{
														TypeName: "int",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithInterfaceWithoutMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type interface{}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &InterfaceSpec{
											Methods: []*Field{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithInterfaceWithMethods(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type interface{
		First(first int) error
		Second(second int) error
	}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &InterfaceSpec{
											Methods: []*Field{
												{
													Name:        "First",
													Annotations: []interface{}{},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name:        "first",
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
												{
													Name:        "Second",
													Annotations: []interface{}{},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name:        "second",
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithTypeGroupWithInterfaceWithMethodsAndComments(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type TypeGroupAnnotation struct {
		Data string
	}

	type TypeAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type FirstArgumentAnnotation struct {
		Data string
	}

	content := `package packageName
// TypeGroupComment
// @TypeGroup({"Data": "TypeGroup"})
type (
	// TypeComment
	// @Type({"Data": "Type"})
	Type interface{
		// FirstComment
		// @First({"Data": "First"})
		First(
			// FirstArgumentComment
			// @FirstArgument({"Data": "FirstArgument"})
			first int,
		) error
	}
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						TypeGroups: []*TypeGroup{
							{
								Comment: `TypeGroupComment
@TypeGroup({"Data": "TypeGroup"})`,
								Annotations: []interface{}{TypeGroupAnnotation{Data: "TypeGroup"}},
								Types: []*Type{
									{
										Name: "Type",
										Comment: `TypeComment
@Type({"Data": "Type"})`,
										Annotations: []interface{}{TypeAnnotation{Data: "Type"}},
										Spec: &InterfaceSpec{
											Methods: []*Field{
												{
													Name: "First",
													Comment: `FirstComment
@First({"Data": "First"})`,
													Annotations: []interface{}{FirstAnnotation{Data: "First"}},
													Spec: &FuncSpec{
														Params: []*Field{
															{
																Name: "first",
																Comment: `FirstArgumentComment
@FirstArgument({"Data": "FirstArgument"})`,
																Annotations: []interface{}{
																	FirstArgumentAnnotation{Data: "FirstArgument"},
																},
																Spec: &SimpleSpec{
																	TypeName: "int",
																},
															},
														},
														Results: []*Field{
															{
																Annotations: []interface{}{},
																Spec: &SimpleSpec{
																	TypeName: "error",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("TypeGroup", TypeGroupAnnotation{}).
		AddAnnotation("Type", TypeAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("FirstArgument", FirstArgumentAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndWithoutArgumentsAndWithoutResultsAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func()
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Spec: &FuncSpec{
									Params:  []*Field{},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncAndArgumentsAndResultsAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func(first firstType, second secondType) (third thirdType, fourth fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},

								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "firstType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "second",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{
										{
											Name:        "third",
											Spec:        &SimpleSpec{TypeName: "thirdType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "fourth",
											Spec:        &SimpleSpec{TypeName: "fourthType"},
											Annotations: []interface{}{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndWithoutArgumentsAndResultsWithoutNamesAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func() (thirdType, fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},

								Spec: &FuncSpec{
									Params: []*Field{},
									Results: []*Field{
										{
											Spec:        &SimpleSpec{TypeName: "thirdType"},
											Annotations: []interface{}{},
										},
										{
											Spec:        &SimpleSpec{TypeName: "fourthType"},
											Annotations: []interface{}{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncAndArgumentsSameTypeAndWithoutResultsAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func(first, second secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},

								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "second",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndArgumentsEllipsisAndWithoutResultsAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func(first firstType, seconds ...secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},

								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "firstType"},
											Annotations: []interface{}{},
										},
										{
											Name: "seconds",
											Spec: &ArraySpec{
												Value:      &SimpleSpec{TypeName: "secondType"},
												IsEllipsis: true,
											},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndFieldsAnnotationsAndWithoutRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	type ThirdAnnotation struct {
		Data string
	}

	type FourthAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func Func(
	// FirstComment
	// @First({"Data": "First"})
	first firstType,

	// SecondComment
	// @Second({"Data": "Second"})
	second secondType,
) (
	// ThirdComment
	// @Third({"Data": "Third"})
	third thirdType,

	// FourthComment
	// @Fourth({"Data": "Fourth"})
	fourth fourthType,
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "first",
											Comment: `FirstComment
@First({"Data": "First"})`,
											Annotations: []interface{}{FirstAnnotation{Data: "First"}},
											Spec:        &SimpleSpec{TypeName: "firstType"},
										},
										{
											Name: "second",
											Comment: `SecondComment
@Second({"Data": "Second"})`,
											Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
											Spec:        &SimpleSpec{TypeName: "secondType"},
										},
									},
									Results: []*Field{
										{
											Name: "third",
											Comment: `ThirdComment
@Third({"Data": "Third"})`,
											Annotations: []interface{}{ThirdAnnotation{Data: "Third"}},
											Spec:        &SimpleSpec{TypeName: "thirdType"},
										},
										{
											Name: "fourth",
											Comment: `FourthComment
@Fourth({"Data": "Fourth"})`,
											Annotations: []interface{}{FourthAnnotation{Data: "Fourth"}},
											Spec:        &SimpleSpec{TypeName: "fourthType"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		AddAnnotation("Third", ThirdAnnotation{}).
		AddAnnotation("Fourth", FourthAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndWithoutArgumentsAndWithoutResultsAndWithRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (*relatedType) Func()
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Annotations: []interface{}{},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: true,
									},
								},
								Spec: &FuncSpec{
									Params:  []*Field{},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncAndArgumentsAndResultsAndRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (related *relatedType) Func(first firstType, second secondType) (third thirdType, fourth fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Name:        "related",
									Annotations: []interface{}{},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: true,
									},
								},
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "firstType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "second",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{
										{
											Name:        "third",
											Spec:        &SimpleSpec{TypeName: "thirdType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "fourth",
											Spec:        &SimpleSpec{TypeName: "fourthType"},
											Annotations: []interface{}{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndWithoutArgumentsAndResultsWithoutNamesAndWithRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (*relatedType) Func() (thirdType, fourthType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Annotations: []interface{}{},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: true,
									},
								},
								Spec: &FuncSpec{
									Params: []*Field{},
									Results: []*Field{
										{
											Spec:        &SimpleSpec{TypeName: "thirdType"},
											Annotations: []interface{}{},
										},
										{
											Spec:        &SimpleSpec{TypeName: "fourthType"},
											Annotations: []interface{}{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncAndArgumentsSameTypeAndWithoutResultsAndWithRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (related relatedType) Func(first, second secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Name:        "related",
									Annotations: []interface{}{},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: false,
									},
								},
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
										{
											Name:        "second",
											Spec:        &SimpleSpec{TypeName: "secondType"},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndArgumentsEllipsisAndWithoutResultsAndWithRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (related *relatedType) Func(first firstType, seconds ...secondType)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Name:        "related",
									Annotations: []interface{}{},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: true,
									},
								},
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name:        "first",
											Spec:        &SimpleSpec{TypeName: "firstType"},
											Annotations: []interface{}{},
										},
										{
											Name: "seconds",
											Spec: &ArraySpec{
												Value:      &SimpleSpec{TypeName: "secondType"},
												IsEllipsis: true,
											},
											Annotations: []interface{}{},
										},
									},
									Results: []*Field{},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithFuncTypeAndFieldsAnnotationsAndWithRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type FuncAnnotation struct {
		Data string
	}

	type RelatedAnnotation struct {
		Data string
	}

	type FirstAnnotation struct {
		Data string
	}

	type SecondAnnotation struct {
		Data string
	}

	type ThirdAnnotation struct {
		Data string
	}

	type FourthAnnotation struct {
		Data string
	}

	content := `package packageName
// FuncComment
// @Func({"Data": "Func"})
func (
	// RelatedComment
	// @Related({"Data": "Related"})
	related *relatedType,
) Func(
	// FirstComment
	// @First({"Data": "First"})
	first firstType,

	// SecondComment
	// @Second({"Data": "Second"})
	second secondType,
) (
	// ThirdComment
	// @Third({"Data": "Third"})
	third thirdType,

	// FourthComment
	// @Fourth({"Data": "Fourth"})
	fourth fourthType,
)
`

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: content,
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     content,
						PackageName: "packageName",
						Annotations: []interface{}{},
						Funcs: []*Func{
							{
								Name: "Func",
								Comment: `FuncComment
@Func({"Data": "Func"})`,
								Annotations: []interface{}{FuncAnnotation{Data: "Func"}},
								Related: &Field{
									Name: "related",
									Comment: `RelatedComment
@Related({"Data": "Related"})`,
									Annotations: []interface{}{
										RelatedAnnotation{Data: "Related"},
									},
									Spec: &SimpleSpec{
										TypeName:  "relatedType",
										IsPointer: true,
									},
								},
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "first",
											Comment: `FirstComment
@First({"Data": "First"})`,
											Annotations: []interface{}{FirstAnnotation{Data: "First"}},
											Spec:        &SimpleSpec{TypeName: "firstType"},
										},
										{
											Name: "second",
											Comment: `SecondComment
@Second({"Data": "Second"})`,
											Annotations: []interface{}{SecondAnnotation{Data: "Second"}},
											Spec:        &SimpleSpec{TypeName: "secondType"},
										},
									},
									Results: []*Field{
										{
											Name: "third",
											Comment: `ThirdComment
@Third({"Data": "Third"})`,
											Annotations: []interface{}{ThirdAnnotation{Data: "Third"}},
											Spec:        &SimpleSpec{TypeName: "thirdType"},
										},
										{
											Name: "fourth",
											Comment: `FourthComment
@Fourth({"Data": "Fourth"})`,
											Annotations: []interface{}{FourthAnnotation{Data: "Fourth"}},
											Spec:        &SimpleSpec{TypeName: "fourthType"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewParser().
		AddAnnotation("Func", FuncAnnotation{}).
		AddAnnotation("Related", RelatedAnnotation{}).
		AddAnnotation("First", FirstAnnotation{}).
		AddAnnotation("Second", SecondAnnotation{}).
		AddAnnotation("Third", ThirdAnnotation{}).
		AddAnnotation("Fourth", FourthAnnotation{}).
		Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestParser_Process_WithInvalidSources(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name:    "file.go",
						Content: "invalid",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(NewParser().Process, storage).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestParser_Process_WithInvalidAnnotation(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type PackageAnnotation struct {
		Data string
	}

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name: "file.go",
						Content: `// PackageComment
// @Package(INVALID)
package packageName
`,
					},
				},
			},
		},
	}

	parser := NewParser().
		AddAnnotation("Package", PackageAnnotation{})

	ctrl.Subtest("WithPanic").
		Call(parser.Process, storage).
		ExpectPanic(ctrl.Type(&json.SyntaxError{}))
}

func TestParser_Process_WithUnknownAnnotation(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "path",
				Files: []*File{
					{
						Name: "file.go",
						Content: `// PackageComment
// @Unknown()
package packageName
`,
					},
				},
			},
		},
	}

	ctrl.Subtest("WithPanic").
		Call(NewParser().Process, storage).
		ExpectPanic(NewErrorf("Unknown annotation '%s'", "Unknown"))
}
