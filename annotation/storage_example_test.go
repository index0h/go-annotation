package annotation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ExampleStorage_ScanRecursive() {
	rootPath, err := ioutil.TempDir("", "example-storage")
	defer func() {
		if err := os.RemoveAll(rootPath); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(rootPath, "child"), 0777); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(rootPath, "root.go"), []byte("package root"), 0666); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(rootPath, "child", "child.go"), []byte("package child"), 0666); err != nil {
		panic(err)
	}

	storage := NewStorage()
	storage.ScanRecursive("namespace/root", rootPath)

	fmt.Printf("Namespace[0].Name:             %s\n", storage.Namespaces[0].Name)
	fmt.Printf("Namespace[0].Files[0].Content: %s\n", storage.Namespaces[0].Files[0].Content)
	fmt.Printf("Namespace[1].Name:             %s\n", storage.Namespaces[1].Name)
	fmt.Printf("Namespace[1].Files[0].Content: %s\n", storage.Namespaces[1].Files[0].Content)

	// Output:
	// Namespace[0].Name:             namespace/root
	// Namespace[0].Files[0].Content: package root
	// Namespace[1].Name:             namespace/root/child
	// Namespace[1].Files[0].Content: package child
}

func ExampleStorage_ScanRecursive_ignore() {
	rootPath, err := ioutil.TempDir("", "example-storage")
	defer func() {
		if err := os.RemoveAll(rootPath); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(rootPath, "child"), 0777); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(rootPath, "root.go"), []byte("package root"), 0666); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(rootPath, "child", "child.go"), []byte("package child"), 0666); err != nil {
		panic(err)
	}

	storage := NewStorage()
	storage.ScanRecursive("namespace/root", rootPath, "child")

	fmt.Printf("Namespace[0].Name:             %s\n", storage.Namespaces[0].Name)
	fmt.Printf("Namespace[0].Files[0].Content: %s\n", storage.Namespaces[0].Files[0].Content)
	fmt.Printf("Namespace[1].Name:             %s\n", storage.Namespaces[1].Name)
	fmt.Printf("Namespace[1].IsIgnored:        %t\n", storage.Namespaces[1].IsIgnored)
	fmt.Printf("len(Namespace[1].Files):       %d\n", len(storage.Namespaces[1].Files))

	// Output:
	// Namespace[0].Name:             namespace/root
	// Namespace[0].Files[0].Content: package root
	// Namespace[1].Name:             namespace/root/child
	// Namespace[1].IsIgnored:        true
	// len(Namespace[1].Files):       0
}
