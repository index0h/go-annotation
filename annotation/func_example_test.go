package annotation

import (
	"fmt"
)

func ExampleFunc_String() {
	model := &Func{
		Name:    "Sum",
		Content: "return x + y",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "x",
					Spec: &SimpleSpec{
						TypeName: "int",
					},
				},
				{
					Name: "y",
					Spec: &SimpleSpec{
						TypeName: "int",
					},
				},
			},
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "int",
					},
				},
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// func Sum(x int, y int) (int) {
	// return x + y
	// }
}

func ExampleFunc_String_related() {
	model := &Func{
		Name:    "ToJSON",
		Content: "result, _ := json.Marshal(r)\nreturn result",
		Related: &Field{
			Name: "r",
			Spec: &SimpleSpec{
				TypeName:  "Related",
				IsPointer: true,
			},
		},
		Spec: &FuncSpec{
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// func (r *Related) ToJSON() (string) {
	// result, _ := json.Marshal(r)
	// return result
	// }
}

func ExampleFunc_String_variadic() {
	model := &Func{
		Name:    "Join",
		Content: "return strings.Join(\", \", parts...)",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "parts",
					Spec: &ArraySpec{
						Value: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
			Results: []*Field{
				{
					Name: "result",
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
			IsVariadic: true,
		},
	}

	fmt.Println(model.String())

	// Output:
	// func Join(parts ...string) (result string) {
	// return strings.Join(", ", parts...)
	// }
}

func ExampleFunc_FetchImports() {
	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{Namespace: "json"},
				},
			},
		},
	}

	model := &Func{
		Name:    "ToJSON",
		Content: "result, _ := json.Marshal(r)\nreturn result",
		Related: &Field{
			Name: "r",
			Spec: &SimpleSpec{
				TypeName:  "Related",
				IsPointer: true,
			},
		},
		Spec: &FuncSpec{
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "json"
}

func ExampleFunc_RenameImports() {
	model := &Func{
		Name:    "ToJSON",
		Content: "result, _ := json.Marshal(r)\nreturn result",
		Related: &Field{
			Name: "r",
			Spec: &SimpleSpec{
				TypeName:  "Related",
				IsPointer: true,
			},
		},
		Spec: &FuncSpec{
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
		},
	}

	model.RenameImports("json", "custom_json")

	fmt.Println(model.String())

	// Output:
	// func (r *Related) ToJSON() (string) {
	// result, _ := custom_json.Marshal(r)
	// return result
	// }
}
