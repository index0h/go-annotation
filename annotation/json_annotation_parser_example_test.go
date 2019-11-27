package annotation

import "fmt"

func ExampleJSONAnnotationParser_Parse() {
	type annotationType struct {
		Data string
	}
	content := `Comment start text
@example({"Data": "Example"})
Comment end text`

	parser := NewJSONAnnotationParser()
	parser.SetAnnotation("example", annotationType{})

	annotations := parser.Parse(content)

	fmt.Printf("%+v", annotations)

	// Output:
	// [{Data:Example}]
}
