package annotation

import "fmt"

func ExampleMapSpec_String() {
	model := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyType",
		},
		Value: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "valueType",
			IsPointer:   true,
		},
	}

	fmt.Println(model.String())

	// Output:
	// map[keyType]*packageName.valueType
}

func ExampleMapSpec_FetchImports() {
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

	model := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "string",
		},
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

func ExampleMapSpec_RenameImports() {
	model := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyType",
		},
		Value: &SimpleSpec{
			PackageName: "strings",
			TypeName:    "Builder",
		},
	}

	model.RenameImports("strings", "custom_strings")

	fmt.Println(model.String())

	// Output:
	// map[keyType]custom_strings.Builder
}
