package annotation

import (
	"fmt"
)

func ExampleType_String() {
	model := &Type{
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
	}

	fmt.Println(model.String())

	// Output:
	// // Stringer interface
	// type Stringer interface{
	// String() (string)
	// }
}

func ExampleType_FetchImports() {
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

	model := &Type{
		Name: "OneHour",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "time"
}

func ExampleType_RenameImports() {
	model := &Type{
		Name: "OneHour",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
	}

	model.RenameImports("time", "custom_time")

	fmt.Println(model.String())

	// Output:
	// type OneHour custom_time.Duration
}
