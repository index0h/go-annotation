package annotation

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type GoScanner struct {
	sourceParser     SourceParser
	annotationParser AnnotationParser
}

func NewGoScanner(sourceParser SourceParser, annotationParser AnnotationParser) *GoScanner {
	if sourceParser == nil {
		panic(errors.New("Variable 'sourceParser' must be not nil"))
	}

	if annotationParser == nil {
		panic(errors.New("Variable 'annotationParser' must be not nil"))
	}

	return &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}
}

// Scans all golang sources recursively inside of rootPath argument.
// If rootNamespace is empty rootPath will be ignored, and Namespace models will be created only for children folders.
// Argument may contain part of path, or absolute path to folder, which must be ignored.
func (s *GoScanner) Scan(storage *Storage, rootNamespace string, rootPath string, ignores ...string) {
	for _, folder := range s.findAllFolders(rootPath) {
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

		namespace.Files = s.scanFiles(namespace.Path)

		storage.Namespaces = append(storage.Namespaces, namespace)
	}
}

// Creates list of File models by *.go files stored in path argument.
func (s *GoScanner) scanFiles(path string) []*File {
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

		result = append(result, s.sourceParser.Parse(file.Name(), string(content)))
	}

	return result
}

func (s *GoScanner) findAllFolders(path string) []string {
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
