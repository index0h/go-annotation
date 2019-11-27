package annotation

import "fmt"

func ExampleArraySpec_String_slice() {
	model := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "string",
		},
	}

	fmt.Println(model.String())

	// Output:
	// []string
}

func ExampleArraySpec_String_array() {
	model := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
			IsPointer:   true,
		},
		Length: "5",
	}

	fmt.Println(model.String())

	// Output:
	// [5]*packageName.typeName
}

func ExampleArraySpec_FetchImports() {
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

	model := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: "strings",
			TypeName:    "Builder",
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "strings"
}

func ExampleArraySpec_RenameImports() {
	model := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: "strings",
			TypeName:    "Builder",
		},
	}

	model.RenameImports("strings", "custom_strings")

	fmt.Println(model.String())

	// Output:
	// []custom_strings.Builder
}

func ExampleArraySpec_RenameImports_length() {
	model := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "string",
		},
		Length: "packageName.Length",
	}

	model.RenameImports("packageName", "custom_packageName")

	fmt.Println(model.String())

	// Output:
	// [custom_packageName.Length]string
}
