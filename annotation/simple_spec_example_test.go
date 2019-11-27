package annotation

import "fmt"

func ExampleSimpleSpec_String() {
	model := &SimpleSpec{
		TypeName: "string",
	}

	fmt.Println(model.String())

	// Output:
	// string
}

func ExampleSimpleSpec_String_pointer() {
	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	fmt.Println(model.String())

	// Output:
	// *packageName.typeName
}

func ExampleSimpleSpec_FetchImports() {
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

	model := &SimpleSpec{
		PackageName: "strings",
		TypeName:    "Builder",
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "strings"
}

func ExampleSimpleSpec_RenameImports() {
	model := &SimpleSpec{
		PackageName: "strings",
		TypeName:    "Builder",
	}

	model.RenameImports("strings", "custom_strings")

	fmt.Println(model.String())

	// Output:
	// custom_strings.Builder
}
