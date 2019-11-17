package annotation

import (
	"io/ioutil"
	"path/filepath"
)

type Scanner struct {
	namespace    string
	path         string
	ignoredPaths map[string]bool
}

func NewScanner(namespace string, path string, ignoredPaths ...string) *Scanner {
	result := &Scanner{
		namespace:    namespace,
		path:         path,
		ignoredPaths: make(map[string]bool, len(ignoredPaths)),
	}

	for _, ignoredPath := range ignoredPaths {
		if filepath.IsAbs(ignoredPath) {
			if ignoredPath, err := filepath.Abs(filepath.Join(path, ignoredPath)); err != nil {
				panic(err)
			} else {
				result.ignoredPaths[ignoredPath] = true
			}
		} else {
			result.ignoredPaths[ignoredPath] = true
		}

	}

	return result
}

func (s *Scanner) Process(storage *Storage) {
	namespaces := []*Namespace{
		{
			Name: s.namespace,
			Path: s.path,
		},
	}

	for i := 0; i < len(namespaces); i++ {
		namespaceEntity := namespaces[i]
		files, err := ioutil.ReadDir(namespaceEntity.Path)

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			path := filepath.Join(namespaceEntity.Path, file.Name())

			if file.IsDir() {
				if _, ok := s.ignoredPaths[path]; ok {
					continue
				}

				addNamespaceEntity := &Namespace{
					Name:  namespaceEntity.Name + "/" + file.Name(),
					Path:  path,
					Files: []*File{},
				}

				namespaces = append(namespaces, addNamespaceEntity)

				continue
			}

			if filepath.Ext(file.Name()) != ".go" {
				continue
			}

			content, err := ioutil.ReadFile(path)

			if err != nil {
				panic(err)
			}

			fileEntity := &File{
				Name:    file.Name(),
				Content: string(content),
			}

			namespaceEntity.Files = append(namespaceEntity.Files, fileEntity)
		}

		if len(namespaceEntity.Files) > 0 {
			if storage.Namespaces == nil {
				storage.Namespaces = []*Namespace{}
			}

			storage.Namespaces = append(storage.Namespaces, namespaceEntity)
		}
	}
}
