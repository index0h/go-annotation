package annotation

import (
	"fmt"
)

func ExampleImportGroup_String() {
	model := &ImportGroup{
		Comment: "Some imports",
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "vendor/project/namespace",
				Comment:   "Import comment",
			},
			{
				Namespace: "vendor/project/namespace/another",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// // Some imports
	// import (
	// // Import comment
	// alias "vendor/project/namespace"
	// "vendor/project/namespace/another"
	// )
}

func ExampleImportGroup_String_oneImport() {
	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "vendor/project/namespace",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// import alias "vendor/project/namespace"
}

func ExampleImportGroup_RenameImports() {
	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "vendor/project/namespace",
				Comment:   "Import comment",
			},
			{
				Namespace: "vendor/project/namespace/another",
			},
		},
	}

	model.RenameImports("another", "custom_another")

	fmt.Println(model.String())

	// Output:
	// import (
	// // Import comment
	// alias "vendor/project/namespace"
	// custom_another "vendor/project/namespace/another"
	// )
}
