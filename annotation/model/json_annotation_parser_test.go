package model

import (
	"encoding/json"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewJSONAnnotationParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
		},
	}

	actual := NewJSONAnnotationParser()

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_SetAnnotation(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "simpleAnnotation"
	annotation := &SimpleSpec{
		TypeName: "value",
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
		},
	}

	expectedParser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
			name:              annotation,
		},
	}

	parser.SetAnnotation(name, annotation)

	ctrl.AssertEqual(expectedParser, parser)
	ctrl.AssertSame(annotation, parser.annotations[name])
}

func TestJSONAnnotationParser_SetAnnotation_WithUpdate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "simpleAnnotation"
	annotation1 := &SimpleSpec{
		TypeName: "annotation1",
	}
	annotation2 := &SimpleSpec{
		TypeName: "annotation1",
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
			name:              annotation1,
		},
	}

	expectedParser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
			name:              annotation2,
		},
	}

	parser.SetAnnotation(name, annotation2)

	ctrl.AssertEqual(expectedParser, parser)
	ctrl.AssertSame(annotation2, parser.annotations[name])
}

func TestJSONAnnotationParser_SetAnnotation_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := ""
	value := &SimpleSpec{
		TypeName: "annotation",
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
		},
	}

	ctrl.Subtest("").
		Call(parser.SetAnnotation, name, value).
		ExpectPanic(NewErrorMessageConstraint("Variable 'name' must be not empty"))
}

func TestJSONAnnotationParser_SetAnnotation_WithProtectedAnnotationName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "FileIsGenerated"
	value := &SimpleSpec{
		TypeName: "annotation",
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
		},
	}

	ctrl.Subtest("").
		Call(parser.SetAnnotation, name, value).
		ExpectPanic(NewErrorMessageConstraint("Annotation name 'FileIsGenerated' is not allowed for change"))
}

func TestJSONAnnotationParser_Parse_WithBool(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value bool

	content := `
value() comment

@value()

	value() comment

@value(false)

@value(

)

value() comment

@value(
	true   
)`

	expected := []interface{}{
		value(false),
		value(false),
		value(false),
		value(true),
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value(false),
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithInt(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value int

	content := `
value() comment

@value()

	value() comment

@value(100)

@value(

)

value() comment

@value(
	-100   
)`

	expected := []interface{}{
		value(0),
		value(100),
		value(0),
		value(-100),
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value(0),
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithFloat(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value float32

	content := `
value() comment

@value()

	value() comment

@value(100)

@value(

)

value() comment

@value(
	-100.55   
)`

	expected := []interface{}{
		value(0),
		value(100),
		value(0),
		value(-100.55),
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value(0),
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithString(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value string

	content := `
value() comment

@value()

	value() comment

@value("data1")

@value(

)

value() comment

@value(
	"data2"   
)`

	expected := []interface{}{
		value(""),
		value("data1"),
		value(""),
		value("data2"),
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value(""),
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithArray(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value []string

	content := `
value() comment

@value()

	value() comment

@value(["data1"])

@value(
[
]
)

value() comment

@value(       [
	"data2"   ,
	"data3"   
              ]
)`

	expected := []interface{}{
		value(nil),
		value{"data1"},
		value{},
		value{"data2", "data3"},
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value{},
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithObject(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value struct {
		First  string
		Second int
	}

	content := `
value() comment

@value()

	value() comment

@value({"First":"data1"})

@value(
{
}
)

value() comment

@value(       {
	"Second":   50   ,
	"First":"data2"   
              }
)`

	expected := []interface{}{
		value{},
		value{First: "data1"},
		value{},
		value{First: "data2", Second: 50},
	}

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value{},
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertEqual(expected, actual)
}

func TestJSONAnnotationParser_Parse_WithEmptyContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value struct {
		First  string
		Second int
	}

	content := ``

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value{},
		},
	}

	actual := parser.Parse(content)

	ctrl.AssertNil(actual)
}

func TestJSONAnnotationParser_Parse_WithUnknownAnnotation(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	content := `@value()`

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{},
	}

	ctrl.Subtest("").
		Call(parser.Parse, content).
		ExpectPanic(NewErrorMessageConstraint("Unknown annotation name 'value'"))
}

func TestJSONAnnotationParser_Parse_WithParseError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "value"
	type value bool

	content := `@value([])`

	parser := &JSONAnnotationParser{
		annotations: map[string]interface{}{
			name: value(false),
		},
	}

	ctrl.Subtest("").
		Call(parser.Parse, content).
		ExpectPanic(ctrl.Type(&json.UnmarshalTypeError{}))
}
