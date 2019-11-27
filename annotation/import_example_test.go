package annotation

import (
	"fmt"
)

func ExampleImport_String() {
	model := &Import{
		Alias:     "alias",
		Namespace: "vendor/project/namespace",
		Comment:   "Import comment",
	}

	fmt.Println(model.String())

	// Output:
	// // Import comment
	// import alias "vendor/project/namespace"
}

func ExampleImport_String_withoutAlias() {
	model := &Import{
		Namespace: "vendor/project/namespace",
		Comment:   "Import comment",
	}

	fmt.Println(model.String())

	// Output:
	// // Import comment
	// import "vendor/project/namespace"
}

func ExampleImport_RenameImports() {
	model := &Import{
		Alias:     "alias",
		Namespace: "vendor/project/namespace",
	}

	model.RenameImports("alias", "custom_alias")

	fmt.Println(model.String())

	// Output:
	// import custom_alias "vendor/project/namespace"
}
