package annotation

import "fmt"

func ExampleEntityRenderer_Render_simpleSpec() {
	entity := &SimpleSpec{
		TypeName: "string",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// string
}

func ExampleEntityRenderer_Render_simpleSpecPointer() {
	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// *packageName.typeName
}

func ExampleEntityRenderer_Render_arraySpecSlice() {
	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "string",
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// []string
}

func ExampleEntityRenderer_Render_arraySpecArray() {
	entity := &ArraySpec{
		Value: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
			IsPointer:   true,
		},
		Length: "5",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// [5]*packageName.typeName
}

func ExampleEntityRenderer_Render_mapSpec() {
	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyType",
		},
		Value: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "valueType",
			IsPointer:   true,
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// map[keyType]*packageName.valueType
}

func ExampleEntityRenderer_Render_funcSpec() {
	entity := &FuncSpec{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// (x int, y int) (int)
}

func ExampleEntityRenderer_Render_interfaceSpec() {
	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Comment: "Include another interface",
				Spec: &SimpleSpec{
					PackageName: "fmt",
					TypeName:    "Stringer",
				},
			},
			{
				Name: "Clone",
				Spec: &FuncSpec{
					Results: []*Field{
						{
							Spec: &InterfaceSpec{},
						},
					},
				},
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// interface{
	// // Include another interface
	// fmt.Stringer
	// Clone() (interface{})
	// }
}

func ExampleEntityRenderer_Render_structSpec() {
	entity := &StructSpec{
		Fields: []*Field{
			{
				Comment: "Include another struct",
				Tag:     "builderTag",
				Spec: &SimpleSpec{
					PackageName: "strings",
					TypeName:    "Builder",
				},
			},
			{
				Name: "Clone",
				Spec: &FuncSpec{
					Results: []*Field{
						{
							Spec: &StructSpec{},
						},
					},
				},
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// struct{
	// // Include another struct
	// strings.Builder "builderTag"
	// Clone func () (struct{})
	// }
}

func ExampleEntityRenderer_Render_import() {
	entity := &Import{
		Alias:     "alias",
		Namespace: "vendor/project/namespace",
		Comment:   "Import comment",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Import comment
	// import alias "vendor/project/namespace"
}

func ExampleEntityRenderer_Render_importWithoutAlias() {
	entity := &Import{
		Namespace: "vendor/project/namespace",
		Comment:   "Import comment",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Import comment
	// import "vendor/project/namespace"
}

func ExampleEntityRenderer_Render_importGroup() {
	entity := &ImportGroup{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Some imports
	// import (
	// // Import comment
	// alias "vendor/project/namespace"
	// "vendor/project/namespace/another"
	// )
}

func ExampleEntityRenderer_Render_importGroupOneImport() {
	entity := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "vendor/project/namespace",
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// import alias "vendor/project/namespace"
}

func ExampleEntityRenderer_Render_const() {
	entity := &Const{
		Name:    "SecondsInDay",
		Value:   "24*60*60",
		Comment: "Count of seconds in one day",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Count of seconds in one day
	// const SecondsInDay = 24*60*60
}

func ExampleEntityRenderer_Render_constSpec() {
	entity := &Const{
		Name:    "SecondsInDay",
		Value:   "24*60*60*time.Second",
		Comment: "Count of seconds in one day",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Count of seconds in one day
	// const SecondsInDay time.Duration = 24*60*60*time.Second
}

func ExampleEntityRenderer_Render_constGroup() {
	entity := &ConstGroup{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Some seconds constants
	// const (
	// // Count of seconds in one day
	// SecondsInDay = 24*60*60
	// SecondsInWeek = SecondsInDay*7
	// CopyOfPreviousConstant
	// )
}

func ExampleEntityRenderer_Render_constGroupOneConst() {
	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "SecondsInDay",
				Value: "24*60*60",
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// const SecondsInDay = 24*60*60
}

func ExampleEntityRenderer_Render_var() {
	entity := &Var{
		Name:    "SecondsInDay",
		Comment: "Count of seconds in one day",
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Count of seconds in one day
	// var SecondsInDay
}

func ExampleEntityRenderer_Render_varSpec() {
	entity := &Var{
		Name:    "SecondsInDay",
		Value:   "24*60*60*time.Second",
		Comment: "Count of seconds in one day",
		Spec: &SimpleSpec{
			PackageName: "time",
			TypeName:    "Duration",
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Count of seconds in one day
	// var SecondsInDay time.Duration = 24*60*60*time.Second
}

func ExampleEntityRenderer_Render_varGroup() {
	entity := &VarGroup{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Some seconds variables
	// var (
	// // Count of seconds in one day
	// SecondsInDay = 24*60*60
	// SecondsInWeek = SecondsInDay*7
	// )
}

func ExampleEntityRenderer_Render_varGroupOneVar() {
	entity := &VarGroup{
		Vars: []*Var{
			{
				Name:  "SecondsInDay",
				Value: "24*60*60",
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// var SecondsInDay = 24*60*60
}

func ExampleEntityRenderer_Render_type() {
	entity := &Type{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Stringer interface
	// type Stringer interface{
	// String() (string)
	// }
}

func ExampleEntityRenderer_Render_typeGroup() {
	entity := &TypeGroup{
		Comment: "Some types",
		Types: []*Type{
			{
				Name: "Duration",
				Spec: &SimpleSpec{
					PackageName: "time",
					TypeName:    "Duration",
				},
			},
			{
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
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Some types
	// type (
	// Duration time.Duration
	// // Stringer interface
	// Stringer interface{
	// String() (string)
	// }
	// )
}

func ExampleEntityRenderer_Render_typeGroupOneType() {
	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "Stringer",
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
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// type Stringer interface{
	// String() (string)
	// }
}

func ExampleEntityRenderer_Render_func() {
	entity := &Func{
		Name:    "Sum",
		Content: "return x + y",
		Spec: &FuncSpec{
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
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// func Sum(x int, y int) (int) {
	// return x + y
	// }
}

func ExampleEntityRenderer_Render_funcRelated() {
	entity := &Func{
		Name:    "ToJSON",
		Content: "result, _ := json.Marshal(r)\nreturn result",
		Related: &Field{
			Name: "r",
			Spec: &SimpleSpec{
				TypeName:  "Related",
				IsPointer: true,
			},
		},
		Spec: &FuncSpec{
			Results: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// func (r *Related) ToJSON() (string) {
	// result, _ := json.Marshal(r)
	// return result
	// }
}

func ExampleEntityRenderer_Render_funcVariadic() {
	entity := &Func{
		Name:    "Join",
		Content: "return strings.Join(\", \", parts...)",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Name: "parts",
					Spec: &ArraySpec{
						Value: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
			Results: []*Field{
				{
					Name: "result",
					Spec: &SimpleSpec{
						TypeName: "string",
					},
				},
			},
			IsVariadic: true,
		},
	}

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// func Join(parts ...string) (result string) {
	// return strings.Join(", ", parts...)
	// }
}

func ExampleEntityRenderer_Render_file() {
	entity := &File{
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

	fmt.Println(NewEntityRenderer().Render(entity))

	// Output:
	// // Hello world
	// package main
	//
	// import "fmt"
	//
	// func main() {
	// 	fmt.Println("Hello world")
	// }
}
