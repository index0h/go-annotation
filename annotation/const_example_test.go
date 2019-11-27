package annotation

import (
	"fmt"
)

func ExampleConst_String() {
	model := &Const{
		Name:    "SecondsInDay",
		Value:   "24*60*60",
		Comment: "Count of seconds in one day",
	}

	fmt.Println(model.String())

	// Output:
	// // Count of seconds in one day
	// const SecondsInDay = 24*60*60
}

func ExampleConst_String_spec() {
	model := &Const{
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
	// const SecondsInDay time.Duration = 24*60*60*time.Second
}

func ExampleConst_FetchImports() {
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

	model := &Const{
		Name: "OneHour",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
		Value: "60*60*time.Second",
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "time"
}

func ExampleConst_FetchImports_value() {
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

	model := &Const{
		Name:  "OneHour",
		Value: "60*60*time.Second",
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "time"
}

func ExampleConst_RenameImports() {
	model := &Const{
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
	// const OneHour custom_time.Duration = 60*60*custom_time.Second
}

func ExampleConst_RenameImports_value() {
	model := &Const{
		Name:  "OneHour",
		Value: "60*60*time.Second",
	}

	model.RenameImports("time", "custom_time")

	fmt.Println(model.String())

	// Output:
	// const OneHour = 60*60*custom_time.Second
}
