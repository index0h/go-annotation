package annotation

import (
	"github.com/pkg/errors"
)

type Storage struct {
	Namespaces []*Namespace
}

// Returns namespace be its name.
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
