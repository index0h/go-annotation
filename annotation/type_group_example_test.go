package annotation

import "fmt"

func ExampleTypeGroup_String() {
	model := &TypeGroup{
		Comment: "Some types",
		Types: []*Type{
			{
				Name: "Duration",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
			},
			{
				Name:    "Stringer",
				Comment: "Stringer interface",
				Spec: &InterfaceSpec{
					Fields: []*Field{
						{
							Name: "String",
							Spec: &FuncSpec{
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
	}

	fmt.Println(model.String())

	// Output:
	// // Some types
	// type (
	// Duration time.Duration
	// // Stringer interface
	// Stringer interface{
	// String() (string)
	// }
	// )
}

func ExampleTypeGroup_String_oneType() {
	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "Stringer",
				Spec: &InterfaceSpec{
					Fields: []*Field{
						{
							Name: "String",
							Spec: &FuncSpec{
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
	}

	fmt.Println(model.String())

	// Output:
	// type Stringer interface{
	// String() (string)
	// }
}

func ExampleTypeGroup_FetchImports() {
	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{Namespace: "time"},
					{Namespace: "fmt"},
				},
			},
		},
	}

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "Hour",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
			},
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "time"
}

func ExampleTypeGroup_RenameImports() {
	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "Hour",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
			},
		},
	}

	model.RenameImports("time", "custom_time")

	fmt.Println(model.String())

	// Output:
	// type Hour custom_time.Duration
}
