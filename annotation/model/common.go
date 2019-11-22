package model

import (
	"encoding/json"
	"reflect"
	"regexp"
)

var identRegexp = regexp.MustCompile("^[\\p{L}_][\\p{L}0-9_]*$")

func cloneAnnotations(annotations []interface{}) []interface{} {
	if annotations == nil {
		return nil
	}

	result := make([]interface{}, len(annotations))

	for i, annotation := range annotations {
		data, err := json.Marshal(annotation)

		if err != nil {
			panic(err)
		}

		annotationCopy := reflect.New(reflect.TypeOf(annotation)).Interface()

		if len(data) > 0 {
			if err = json.Unmarshal(data, &annotationCopy); err != nil {
				panic(err)
			}
		}

		result[i] = reflect.ValueOf(annotationCopy).Elem().Interface()
	}

	return result
}

func uniqImports(all []*Import) []*Import {
	if all == nil {
		return nil
	}

	result := []*Import{}

	for _, element := range all {
		isUniq := true

		for _, resultElement := range result {
			if resultElement == element || (resultElement.Alias == element.Alias &&
				resultElement.Namespace == element.Namespace &&
				resultElement.Comment == element.Comment) {
				isUniq = false

				break
			}
		}

		if isUniq {
			result = append(result, element)
		}
	}

	return result
}