package model

import (
	"github.com/pkg/errors"
)

type Storage struct {
	Namespaces []*Namespace
}

func NewStorage() *Storage {
	return &Storage{
		Namespaces: []*Namespace{},
	}
}

func (m *Storage) Validate() {
	namespaceNames := map[string]bool{}
	namespacePaths := map[string]bool{}

	for i, element := range m.Namespaces {
		if element == nil {
			panic(errors.Errorf("Variable 'Namespaces[%d]' must be not nil", i))
		}

		element.Validate()

		if _, ok := namespaceNames[element.Name]; ok {
			panic(errors.Errorf("Storage has duplicate namespace 'Name': '%s'", element.Name))
		} else {
			namespaceNames[element.Name] = true
		}

		if _, ok := namespacePaths[element.Path]; ok {
			panic(errors.Errorf("Storage has duplicate namespace 'Path': '%s'", element.Path))
		} else {
			namespacePaths[element.Path] = true
		}
	}
}

func (m *Storage) Clone() interface{} {
	result := &Storage{}

	if m.Namespaces != nil {
		result.Namespaces = make([]*Namespace, len(m.Namespaces))
	}

	for i, element := range m.Namespaces {
		result.Namespaces[i] = element.Clone().(*Namespace)
	}

	return result
}

func (m *Storage) FindNamespaceByName(name string) *Namespace {
	if name == "" {
		panic(errors.New("Variable 'name' must be not empty"))
	}

	for _, element := range m.Namespaces {
		if name == element.Name {
			return element
		}
	}

	return nil
}
