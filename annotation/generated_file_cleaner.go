package annotation

import (
	"os"
	"path/filepath"
)

type FileIsGeneratedAnnotation bool

type GeneratedFileCleaner struct {
}

func NewGeneratedFileCleaner() *GeneratedFileCleaner {
	return &GeneratedFileCleaner{}
}

// Removes old generated files with content and FileIsGeneratedAnnotation annotation.
func (*GeneratedFileCleaner) Clean(storage *Storage) {
	for _, namespace := range storage.Namespaces {
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
