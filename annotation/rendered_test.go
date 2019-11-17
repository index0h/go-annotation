package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewRenderer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("WithPositiveResult").
		Call(NewRenderer).
		ExpectResult(&Renderer{})
}

func TestRenderer_Process_WithPackage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     Header + "package packageName\n",
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithPackageAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Comment:     "comment\nhere",
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `// comment
// here
package packageName
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Comment:     "comment\nhere",
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

import alias "namespace"
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneImportAndGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Import
// Group
// Comment
import alias "namespace"
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithImportGroupByImportOwnComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
										Comment:   "Import\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

import (
	// Import
	// Comment
	alias "namespace"
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
										Comment:   "Import\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithImportGroupByImportOwnCommentAndGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
										Comment:   "Import\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Import
// Group
// Comment
import (
	// Import
	// Comment
	alias "namespace"
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias",
										Namespace: "namespace",
										Comment:   "Import\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
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
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

import (
	alias1 "namespace1"
	alias2 "namespace2"
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
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
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithImportGroupAndImportComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias1",
										Namespace: "namespace1",
										Comment:   "Import1\nComment",
									},
									{
										Alias:     "alias2",
										Namespace: "namespace2",
										Comment:   "Import2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

import (
	// Import1
	// Comment
	alias1 "namespace1"
	// Import2
	// Comment
	alias2 "namespace2"
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Imports: []*Import{
									{
										Alias:     "alias1",
										Namespace: "namespace1",
										Comment:   "Import1\nComment",
									},
									{
										Alias:     "alias2",
										Namespace: "namespace2",
										Comment:   "Import2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithImportGroupAndImportCommentAndImportGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias1",
										Namespace: "namespace1",
										Comment:   "Import1\nComment",
									},
									{
										Alias:     "alias2",
										Namespace: "namespace2",
										Comment:   "Import2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Import
// Group
// Comment
import (
	// Import1
	// Comment
	alias1 "namespace1"
	// Import2
	// Comment
	alias2 "namespace2"
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ImportGroups: []*ImportGroup{
							{
								Comment: "Import\nGroup\nComment",
								Imports: []*Import{
									{
										Alias:     "alias1",
										Namespace: "namespace1",
										Comment:   "Import1\nComment",
									},
									{
										Alias:     "alias2",
										Namespace: "namespace2",
										Comment:   "Import2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneConstAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const ConstName = ConstValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneConstAndValueAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const ConstName = ConstValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneConstAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const ConstName SpecPackage.SpecType = ConstValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneConstAndSpecAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const ConstName SpecPackage.SpecType = ConstValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:  "ConstName",
										Value: "ConstValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupByConstCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	// Const
	// Comment
	ConstName = ConstValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupByConstCommentAndValueAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const (
	// Const
	// Comment
	ConstName = ConstValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupByConstCommentAndValueAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	// Const
	// Comment
	ConstName SpecPackage.SpecType = ConstValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupByConstCommentAndValueAndSpecAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const (
	// Const
	// Comment
	ConstName SpecPackage.SpecType = ConstValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "ConstName",
										Value:   "ConstValue",
										Comment: "Const\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "Const1Name",
										Value: "Const1Value",
									},
									{
										Name:  "Const2Name",
										Value: "Const2Value",
									},
									{
										Name: "Const3Name",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	Const1Name = Const1Value
	Const2Name = Const2Value
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "Const1Name",
										Value: "Const1Value",
									},
									{
										Name:  "Const2Name",
										Value: "Const2Value",
									},
									{
										Name: "Const3Name",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValueAndConstComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	// Const1
	// Comment
	Const1Name = Const1Value
	// Const2
	// Comment
	Const2Name = Const2Value
	// Const3
	// Comment
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValueAndConstCommentAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const (
	// Const1
	// Comment
	Const1Name = Const1Value
	// Const2
	// Comment
	Const2Name = Const2Value
	// Const3
	// Comment
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValueAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "Const1Name",
										Value: "Const1Value",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:  "Const2Name",
										Value: "Const2Value",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name: "Const3Name",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	Const1Name Spec1Package.Spec1Type = Const1Value
	Const2Name Spec2Type              = Const2Value
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:  "Const1Name",
										Value: "Const1Value",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:  "Const2Name",
										Value: "Const2Value",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name: "Const3Name",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValueAndSpecAndConstComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

const (
	// Const1
	// Comment
	Const1Name Spec1Package.Spec1Type = Const1Value
	// Const2
	// Comment
	Const2Name Spec2Type = Const2Value
	// Const3
	// Comment
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithConstGroupAndValueAndSpecAndConstCommentAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Const
// Group
// Comment
const (
	// Const1
	// Comment
	Const1Name Spec1Package.Spec1Type = Const1Value
	// Const2
	// Comment
	Const2Name Spec2Type = Const2Value
	// Const3
	// Comment
	Const3Name
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						ConstGroups: []*ConstGroup{
							{
								Comment: "Const\nGroup\nComment",
								Consts: []*Const{
									{
										Name:    "Const1Name",
										Value:   "Const1Value",
										Comment: "Const1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Const2Name",
										Value:   "Const2Value",
										Comment: "Const2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
										},
									},
									{
										Name:    "Const3Name",
										Comment: "Const3\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var VarName = VarValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndValueAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var VarName = VarValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name: "VarName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var VarName SpecPackage.SpecType
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name: "VarName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndSpecAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name: "VarName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var VarName SpecPackage.SpecType
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name: "VarName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var VarName SpecPackage.SpecType = VarValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneVarAndSpecAndValueAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var VarName SpecPackage.SpecType = VarValue
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:  "VarName",
										Value: "VarValue",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var
	// Comment
	VarName = VarValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndValueAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var (
	// Var
	// Comment
	VarName = VarValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var
	// Comment
	VarName SpecPackage.SpecType
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndSpecAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var (
	// Var
	// Comment
	VarName SpecPackage.SpecType
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndValueAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var
	// Comment
	VarName SpecPackage.SpecType = VarValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupByVarCommentAndValueAndSpecAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var (
	// Var
	// Comment
	VarName SpecPackage.SpecType = VarValue
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "VarName",
										Value:   "VarValue",
										Comment: "Var\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "Var1Name",
										Value: "Var1Value",
									},
									{
										Name:  "Var2Name",
										Value: "Var2Value",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	Var1Name = Var1Value
	Var2Name = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "Var1Name",
										Value: "Var1Value",
									},
									{
										Name:  "Var2Name",
										Value: "Var2Value",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var1
	// Comment
	Var1Name = Var1Value
	// Var2
	// Comment
	Var2Name = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var (
	// Var1
	// Comment
	Var1Name = Var1Value
	// Var2
	// Comment
	Var2Name = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValueAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "Var1Name",
										Value: "Var1Value",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:  "Var2Name",
										Value: "Var2Value",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	Var1Name Spec1Package.Spec1Type = Var1Value
	Var2Name Spec2Type              = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:  "Var1Name",
										Value: "Var1Value",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:  "Var2Name",
										Value: "Var2Value",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndSpecAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var1
	// Comment
	Var1Name Spec1Package.Spec1Type
	// Var2
	// Comment
	Var2Name Spec2Type
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name: "Var1Name",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name: "Var2Name",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	Var1Name Spec1Package.Spec1Type
	Var2Name Spec2Type
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name: "Var1Name",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name: "Var2Name",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValueAndSpecAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

var (
	// Var1
	// Comment
	Var1Name Spec1Package.Spec1Type = Var1Value
	// Var2
	// Comment
	Var2Name Spec2Type = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithVarGroupAndValueAndSpecAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Var
// Group
// Comment
var (
	// Var1
	// Comment
	Var1Name Spec1Package.Spec1Type = Var1Value
	// Var2
	// Comment
	Var2Name Spec2Type = Var2Value
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						VarGroups: []*VarGroup{
							{
								Comment: "Var\nGroup\nComment",
								Vars: []*Var{
									{
										Name:    "Var1Name",
										Value:   "Var1Value",
										Comment: "Var1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Var2Name",
										Value:   "Var2Value",
										Comment: "Var2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name: "TypeName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

type TypeName SpecPackage.SpecType
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name: "TypeName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithOneTypeAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Comment: "Type\nGroup\nComment",
								Types: []*Type{
									{
										Name: "TypeName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Type
// Group
// Comment
type TypeName SpecPackage.SpecType
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Comment: "Type\nGroup\nComment",
								Types: []*Type{
									{
										Name: "TypeName",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithTypeGroupByTypeComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name:    "TypeName",
										Comment: "Type\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

type (
	// Type
	// Comment
	TypeName SpecPackage.SpecType
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name:    "TypeName",
										Comment: "Type\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithTypeGroupByTypeCommentAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Comment: "Type\nGroup\nComment",
								Types: []*Type{
									{
										Name:    "TypeName",
										Comment: "Type\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Type
// Group
// Comment
type (
	// Type
	// Comment
	TypeName SpecPackage.SpecType
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Comment: "Type\nGroup\nComment",
								Types: []*Type{
									{
										Name:    "TypeName",
										Comment: "Type\nComment",
										Spec: &SimpleSpec{
											PackageName: "SpecPackage",
											TypeName:    "SpecType",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name: "Type1Name",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name: "Type2Name",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

type (
	Type1Name Spec1Package.Spec1Type
	Type2Name Spec2Type
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name: "Type1Name",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name: "Type2Name",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithTypeGroupAndTypeComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name:    "Type1Name",
										Comment: "Type1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Type2Name",
										Comment: "Type2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

type (
	// Type1
	// Comment
	Type1Name Spec1Package.Spec1Type
	// Type2
	// Comment
	Type2Name Spec2Type
)
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						TypeGroups: []*TypeGroup{
							{
								Types: []*Type{
									{
										Name:    "Type1Name",
										Comment: "Type1\nComment",
										Spec: &SimpleSpec{
											PackageName: "Spec1Package",
											TypeName:    "Spec1Type",
										},
									},
									{
										Name:    "Type2Name",
										Comment: "Type2\nComment",
										Spec: &SimpleSpec{
											TypeName: "Spec2Type",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

func FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFuncAndFuncComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
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

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Func
// Comment
func FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
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

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFuncAndRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name: "related",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

func (related Related) FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name: "related",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFuncAndRelatedAndFuncComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name: "related",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Func
// Comment
func (related Related) FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name: "related",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFuncAndRelatedAndRelatedComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name:    "related",
									Comment: "Related\nComment",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

func (
	// Related
	// Comment
	related Related) FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name:    "related",
									Comment: "Related\nComment",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}

func TestRenderer_Process_WithFuncAndRelatedAndRelatedCommentAndFuncComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name:        "file.go",
						Content:     "",
						PackageName: "packageName",
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name:    "related",
									Comment: "Related\nComment",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expectedStorage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: "/some/path",
				Files: []*File{
					{
						Name: "file.go",
						Content: Header + `package packageName

// Func
// Comment
func (
	// Related
	// Comment
	related Related) FuncName(argument string) string {
	return argument
}
`,
						PackageName: "packageName",
						Annotations: []interface{}{FileIsGeneratedAnnotation(true)},
						Funcs: []*Func{
							{
								Name:    "FuncName",
								Comment: "Func\nComment",
								Content: "return argument",
								Spec: &FuncSpec{
									Params: []*Field{
										{
											Name: "argument",
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
									Results: []*Field{
										{
											Spec: &SimpleSpec{
												TypeName: "string",
											},
										},
									},
								},
								Related: &Field{
									Name:    "related",
									Comment: "Related\nComment",
									Spec: &SimpleSpec{
										TypeName: "Related",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	(&Renderer{}).Process(storage)

	ctrl.AssertEqual(expectedStorage, storage)
}
