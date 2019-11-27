package annotation

import (
	"fmt"
)

func ExampleFuncSpec_String() {
	model := &FuncSpec{
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
	}

	fmt.Println(model.String())

	// Output:
	// (x int, y int) (int)
}

func ExampleFuncSpec_FetchImports() {
	file := &File{
		Name:        "file.go",
		PackageName: "packageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{Namespace: "log"},
				},
			},
		},
	}

	model := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "log",
					TypeName:    "Logger",
					IsPointer:   true,
				},
			},
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "log"
}

func ExampleFuncSpec_RenameImports() {
	model := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					PackageName: "log",
					TypeName:    "Logger",
					IsPointer:   true,
				},
			},
		},
	}

	model.RenameImports("log", "custom_log")

	fmt.Println(model.String())

	// Output:
	// (*custom_log.Logger)
}
