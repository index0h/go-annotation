package annotation

import (
	"fmt"
)

func ExampleFile_String() {
	model := &File{
		Comment:     "Hello world",
		PackageName: "main",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "fmt",
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "main",
				Content: "fmt.Println(\"Hello world\")",
			},
		},
	}

	fmt.Println(model.String())

	// Output:
	// // Generated by github.com/index0h/go-annotation
	// // DO NOT EDIT
	// // @FileIsGenerated(true)
	// // Hello world
	// package main
	//
	// import "fmt"
	//
	// func main() {
	// 	fmt.Println("Hello world")
	// }
}

func ExampleFile_RenameImports() {
	model := &File{
		PackageName: "main",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: "fmt",
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "main",
				Content: "fmt.Println(\"Hello world\")",
			},
		},
	}

	model.RenameImports("fmt", "custom_fmt")

	fmt.Println(model.String())

	// Output:
	// // Generated by github.com/index0h/go-annotation
	// // DO NOT EDIT
	// // @FileIsGenerated(true)
	// package main
	//
	// import custom_fmt "fmt"
	//
	// func main() {
	// 	custom_fmt.Println("Hello world")
	// }
}