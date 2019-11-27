package annotation

import "fmt"

func ExampleInterfaceSpec_String() {
	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Comment: "Include another interface",
				Spec: &SimpleSpec{
					PackageName: "fmt",
					TypeName:    "Stringer",
				},
			},
			{
				Name: "Clone",
				Spec: &FuncSpec{
					Results: []*Field{
						{
							Spec: &InterfaceSpec{},
						},
					},
				},
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// interface{
	// // Include another interface
	// fmt.Stringer
	// Clone() (interface{})
	// }
}

func ExampleInterfaceSpec_FetchImports() {
	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{Namespace: "strings"},
					{Namespace: "fmt"},
				},
			},
		},
	}

	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "fmt",
					TypeName:    "Stringer",
				},
			},
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "fmt"
}

func ExampleInterfaceSpec_RenameImports() {
	model := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "fmt",
					TypeName:    "Stringer",
				},
			},
		},
	}

	model.RenameImports("fmt", "custom_fmt")

	fmt.Println(model.String())

	// Output:
	// interface{
	// custom_fmt.Stringer
	// }
}
