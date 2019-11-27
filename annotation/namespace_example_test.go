package annotation

import "fmt"

func ExampleNamespace_PackageName() {
	model := &Namespace{
		Name: "vendor/project/namespace",
		Path: "/path/to/project",
	}

	packageName := model.PackageName()

	fmt.Println(packageName)

	// Output:
	// namespace
}

func ExampleNamespace_PackageName_fromFile() {
	model := &Namespace{
		Name: "vendor/project/namespace",
		Path: "/path/to/project",
		Files: []*File{
			{
				Name:        "file.go",
				PackageName: "packageName",
			},
		},
	}

	packageName := model.PackageName()

	fmt.Println(packageName)

	// Output:
	// packageName
}
