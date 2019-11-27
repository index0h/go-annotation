package annotation

import (
	"fmt"
)

func ExampleVar_String() {
	model := &Var{
		Name:    "SecondsInDay",
		Comment: "Count of seconds in one day",
	}

	fmt.Println(model.String())

	// Output:
	// // Count of seconds in one day
	// var SecondsInDay
}

func ExampleVar_String_spec() {
	model := &Var{
		Name:    "SecondsInDay",
		Value:   "24*60*60*time.Second",
		Comment: "Count of seconds in one day",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
	}

	fmt.Println(model.String())

	// Output:
	// // Count of seconds in one day
	// var SecondsInDay time.Duration = 24*60*60*time.Second
}

func ExampleVar_FetchImports() {
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

	model := &Var{
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

func ExampleVar_RenameImports() {
	model := &Var{
		Name: "OneHour",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
		Value: "60*60*time.Second",
	}

	model.RenameImports("time", "custom_time")

	fmt.Println(model.String())

	// Output:
	// var OneHour custom_time.Duration = 60*60*custom_time.Second
}
