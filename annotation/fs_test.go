package annotation

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/index0h/go-unit/unit"
)

func NewDir(controller *unit.Controller, dir string, name string) string {
	err := os.Mkdir(filepath.Join(dir, name), 0777)

	if err != nil {
		controller.TestingT().Error(err)

		return ""
	}

	controller.RegisterFinish(
		func() {
			if err := os.RemoveAll(dir); err != nil {
				controller.TestingT().Error(err)
			}
		},
	)

	return filepath.Join(dir, name)
}

func NewTempDir(controller *unit.Controller, dir string, prefix string) string {
	result, err := ioutil.TempDir(dir, prefix)

	if err != nil {
		controller.TestingT().Error(err)

		return ""
	}

	controller.RegisterFinish(
		func() {
			if err := os.RemoveAll(result); err != nil {
				controller.TestingT().Error(err)
			}
		},
	)

	return result
}

func NewFile(controller *unit.Controller, dir string, name string) *os.File {
	result, err := os.Create(filepath.Join(dir, name))

	if err != nil {
		controller.TestingT().Error(err)

		return nil
	}

	controller.RegisterFinish(
		func() {
			if err := result.Close(); err != nil {
				controller.TestingT().Error(err)
			}
		},
	)

	return result
}
