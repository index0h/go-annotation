package annotation

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Storage struct {
	AnnotationParser AnnotationParser
	SourceParser     SourceParser
	Namespaces       []*Namespace
}

func NewStorage() *Storage {
	annotationParser := NewJSONAnnotationParser()
	sourceParser := NewGoSourceParser(annotationParser)

	return &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces:       []*Namespace{},
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
	result := &Storage{
		AnnotationParser: m.AnnotationParser,
		SourceParser:     m.SourceParser,
	}

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

func (m *Storage) ScanRecursive(rootNamespace string, rootPath string, ignores ...string) {
	for _, folder := range m.findAllFolders(rootPath) {
		pathSuffix := strings.TrimLeft(folder, rootPath)

		if pathSuffix == "" && rootNamespace == "" {
			continue
		}

		namespace := &Namespace{
			Name: strings.Trim(rootNamespace+"/"+pathSuffix, "/"),
			Path: folder,
		}

		for _, ignore := range ignores {
			if strings.Contains(pathSuffix, ignore) {
				namespace.IsIgnored = true

				break
			}
		}

		if !namespace.IsIgnored {
			namespace.Files = m.ScanFiles(namespace.Path)
		}

		m.Namespaces = append(m.Namespaces, namespace)
	}

	m.Validate()
}

func (m *Storage) ScanFiles(path string) []*File {
	result := []*File{}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())

		if file.IsDir() || filepath.Ext(file.Name()) != ".go" {
			continue
		}

		content, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		result = append(result, m.SourceParser.Parse(file.Name(), string(content)))
	}

	return result
}

func (m *Storage) RemoveOldGeneratedFiles() {
	m.Validate()

	for _, namespace := range m.Namespaces {
		if namespace.IsIgnored {
			continue
		}

		resultFiles := make([]*File, 0, len(namespace.Files))

		for _, file := range namespace.Files {
			removeFile := false

			for _, rawAnnotation := range file.Annotations {
				if annotation, ok := rawAnnotation.(FileIsGeneratedAnnotation); ok && bool(annotation) {
					removeFile = true

					if err := os.Remove(filepath.Join(namespace.Path, file.Name)); err != nil {
						panic(err)
					}

					break
				}
			}

			if !removeFile {
				resultFiles = append(resultFiles, file)
			}
		}

		namespace.Files = resultFiles
	}
}

func (m *Storage) WriteGeneratedFiles() {
	m.Validate()

	for _, namespace := range m.Namespaces {
		if namespace.IsIgnored {
			continue
		}

		for _, file := range namespace.Files {
			if file.Content == "" {
				file.Content = file.String()

				if err := os.MkdirAll(namespace.Path, os.ModePerm); err != nil {
					panic(err)
				}

				filePath := filepath.Join(namespace.Path, file.Name)

				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					panic(errors.Errorf("File '%s' already exists", filePath))
				}

				if err := ioutil.WriteFile(filePath, []byte(file.Content), 0666); err != nil {
					panic(err)
				}
			}
		}
	}
}

func (m *Storage) findAllFolders(path string) []string {
	result := []string{}

	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				result = append(result, path)
			}

			return nil
		},
	)

	if err != nil {
		panic(err)
	}

	sort.Strings(result)

	return result
}
