package annotation

import (
	"fmt"
)

func ExampleConstGroup_String() {
	model := &ConstGroup{
		Comment: "Some seconds constants",
		Consts: []*Const{
			{
				Name:    "SecondsInDay",
				Value:   "24*60*60",
				Comment: "Count of seconds in one day",
			},
			{
				Name:  "SecondsInWeek",
				Value: "SecondsInDay*7",
			},
			{
				Name: "CopyOfPreviousConstant",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// // Some seconds constants
	// const (
	// // Count of seconds in one day
	// SecondsInDay = 24*60*60
	// SecondsInWeek = SecondsInDay*7
	// CopyOfPreviousConstant
	// )
}

func ExampleConstGroup_String_oneConst() {
	model := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "SecondsInDay",
				Value: "24*60*60",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// const SecondsInDay = 24*60*60
}

func ExampleConstGroup_FetchImports() {
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

	model := &ConstGroup{
		Consts: []*Const{
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

func ExampleConstGroup_RenameImports() {
	model := &ConstGroup{
		Consts: []*Const{
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
	// const (
	// OneHour custom_time.Duration = 60*60*custom_time.Second
	// TwoHours custom_time.Duration = 2*60*60*custom_time.Second
	// )
}
