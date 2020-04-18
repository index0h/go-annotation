package annotation

import "fmt"

func ExampleGoSourceParser_Parse() {
	fileName := "file.go"
	content := `// Hello world
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}`

	file := NewGoSourceParser(NewJSONAnnotationParser()).Parse(fileName, content)
	// Clean content to use File rendering
	file.Content = ""

	fmt.Println(NewEntityRenderer().Render(file))

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
