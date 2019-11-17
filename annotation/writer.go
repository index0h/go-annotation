package annotation

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Process(storage *Storage) {
	for _, namespace := range storage.Namespaces {
		for _, file := range namespace.Files {
			for _, rawAnnotation := range file.Annotations {
				if annotation, ok := rawAnnotation.(FileIsGeneratedAnnotation); ok && bool(annotation) {
					if err := os.MkdirAll(namespace.Path, os.ModePerm); err != nil {
						panic(err)
					}

					filePath := filepath.Join(namespace.Path, file.Name)

					if _, err := os.Stat(filePath); !os.IsNotExist(err) {
						panic(NewErrorf("File '%s' already exists", filePath))
					}

					if err := ioutil.WriteFile(filePath, []byte(file.Content), 0664); err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
