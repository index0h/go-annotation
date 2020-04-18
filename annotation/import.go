package annotation

import (
	"path/filepath"
)

// Import represents import declaration.
type Import struct {
	Alias       string
	Namespace   string
	Comment     string
	Annotations []interface{}
}

// Returns Alias field if it's not empty, otherwise base path of Namespace field.
func (m *Import) RealAlias() string {
	if m.Alias != "" {
		return m.Alias
	}

	return filepath.Base(m.Namespace)
}
