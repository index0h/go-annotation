package annotation

import "fmt"

func ExampleStructSpec_String() {
	model := &StructSpec{
		Fields: []*Field{
			{
				Comment: "Include another struct",
				Tag:     "builderTag",
				Spec: &SimpleSpec{
					PackageName: "strings",
					TypeName:    "Builder",
				},
			},
			{
				Name: "Clone",
				Spec: &FuncSpec{
					Results: []*Field{
						{
							Spec: &StructSpec{},
						},
					},
				},
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// struct{
	// // Include another struct
	//  strings.Builder "builderTag"
	// Clone func () (struct{})
	// }
}

func ExampleStructSpec_FetchImports() {
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

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "ToString",
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

func ExampleStructSpec_RenameImports() {
	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "ToString",
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
	// struct{
	// ToString custom_fmt.Stringer
	// }
}
