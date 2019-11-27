package annotation

import (
	"fmt"
)

func ExampleVarGroup_String() {
	model := &VarGroup{
		Comment: "Some seconds variables",
		Vars: []*Var{
			{
				Name:    "SecondsInDay",
				Value:   "24*60*60",
				Comment: "Count of seconds in one day",
			},
			{
				Name:  "SecondsInWeek",
				Value: "SecondsInDay*7",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// // Some seconds variables
	// var (
	// // Count of seconds in one day
	// SecondsInDay = 24*60*60
	// SecondsInWeek = SecondsInDay*7
	// )
}

func ExampleVarGroup_String_oneVar() {
	model := &VarGroup{
		Vars: []*Var{
			{
				Name:  "SecondsInDay",
				Value: "24*60*60",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// var SecondsInDay = 24*60*60
}

func ExampleVarGroup_FetchImports() {
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

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "OneHour",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
				Value: "60*60*time.Second",
			},
		},
	}

	usedImports := model.FetchImports(file)

	fmt.Println(usedImports[0].String())

	// Output:
	// import "time"
}

func ExampleVarGroup_RenameImports() {
	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "OneHour",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
				Value: "60*60*time.Second",
			},
			{
				Name: "TwoHours",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
				Value: "2*60*60*time.Second",
			},
		},
	}

	model.RenameImports("time", "custom_time")

	fmt.Println(model.String())

	// Output:
	// var (
	// OneHour custom_time.Duration = 60*60*custom_time.Second
	// TwoHours custom_time.Duration = 2*60*60*custom_time.Second
	// )
}
