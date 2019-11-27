package annotation

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
)

var jsonAnnotationRegexp = regexp.MustCompile(`(?mU)^@([\p{L}_][\p{L}\d_]*)\(((.|\n)*)\)$`)

var protectedAnnotations = map[string]interface{}{
	"FileIsGenerated": FileIsGeneratedAnnotation(false),
}

// Parsers comment and creates list of annotations.
type JSONAnnotationParser struct {
	annotations map[string]interface{}
}

// Creates new instance of JSONAnnotationParser.
func NewJSONAnnotationParser() *JSONAnnotationParser {
	result := &JSONAnnotationParser{
		annotations: map[string]interface{}{},
	}

	for key, value := range protectedAnnotations {
		result.annotations[key] = value
	}

	return result
}

// Registers new or updates old annotation by name and its type.
func (p *JSONAnnotationParser) SetAnnotation(name string, annotationType interface{}) {
	if name == "" {
		panic(errors.New("Variable 'name' must be not empty"))
	}

	if _, ok := protectedAnnotations[name]; ok {
		panic(errors.Errorf("Annotation name '%s' is not allowed for change", name))
	}

	p.annotations[name] = annotationType
}

// Parsers comment and creates list of annotations.
func (p *JSONAnnotationParser) Parse(content string) []interface{} {
	if content == "" {
		return nil
	}

	result := []interface{}{}

	for _, part := range jsonAnnotationRegexp.FindAllStringSubmatch(content, -1) {
		data := strings.TrimSpace(part[2])
		annotation, ok := p.annotations[part[1]]

		if !ok {
			continue
		}

		value := reflect.New(reflect.TypeOf(annotation)).Interface()

		if len(data) > 0 {
			if err := json.Unmarshal([]byte(data), &value); err != nil {
				panic(err)
			}
		}

		result = append(result, reflect.ValueOf(value).Elem().Interface())
	}

	return result
}
