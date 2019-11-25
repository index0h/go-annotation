package model

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/index0h/go-unit/unit"
)

type TmpFS struct {
	ctrl     *unit.Controller
	rootPath string
}

func NewTmpFS(ctrl *unit.Controller) *TmpFS {
	rootPath, err := ioutil.TempDir("", "")

	if err != nil {
		ctrl.TestingT().Error(err)

		return nil
	}

	result := &TmpFS{
		ctrl:     ctrl,
		rootPath: rootPath,
	}

	ctrl.RegisterFinish(
		func() {
			result.Remove(".")
		},
	)

	return &TmpFS{
		ctrl:     ctrl,
		rootPath: rootPath,
	}
}

func (tmpfs *TmpFS) RootPath() string {
	return tmpfs.rootPath
}

func (tmpfs *TmpFS) Remove(relativePath string) *TmpFS {
	if err := os.RemoveAll(filepath.Join(tmpfs.rootPath, relativePath)); err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	}

	return tmpfs
}

func (tmpfs *TmpFS) CreateFile(relativePath string, fileMode os.FileMode, content string) *TmpFS {
	filePath := filepath.Join(tmpfs.rootPath, relativePath)
	file, err := os.Create(filePath)

	if err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	}

	if err := os.Chmod(filePath, fileMode); err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	}

	if _, err := file.WriteString(content); err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	} else if err = file.Close(); err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	}

	return tmpfs
}

func (tmpfs *TmpFS) CreateDir(relativePath string, fileMode os.FileMode) *TmpFS {
	if err := os.Mkdir(filepath.Join(tmpfs.rootPath, relativePath), fileMode); err != nil {
		tmpfs.ctrl.TestingT().Error(err)
	}

	return tmpfs
}

func (tmpfs *TmpFS) FileExists() unit.Constraint {
	return tmpfs.ctrl.Callback(
		func(path interface{}) bool {
			if _, err := os.Stat(filepath.Join(tmpfs.rootPath, path.(string))); err == nil {
				return true
			} else if os.IsNotExist(err) {
				return false
			} else {
				panic(err)
			}
		},
	)
}

func (tmpfs *TmpFS) FileContent(content string) unit.Constraint {
	return tmpfs.ctrl.Callback(
		func(relativePath interface{}) bool {
			if data, err := ioutil.ReadFile(filepath.Join(tmpfs.rootPath, relativePath.(string))); err != nil {
				return false
			} else {
				return content == string(data)
			}
		},
	)
}

func (tmpfs *TmpFS) AssertFileExists(relativePath string) {
	tmpfs.ctrl.TestingT().Helper()

	tmpfs.ctrl.AssertThat(relativePath, tmpfs.FileExists(), "file path")
}

func (tmpfs *TmpFS) AssertNotFileExists(relativePath string) {
	tmpfs.ctrl.TestingT().Helper()

	tmpfs.ctrl.AssertThat(relativePath, tmpfs.ctrl.Not(tmpfs.FileExists()), "file path")
}

func (tmpfs *TmpFS) AssertFileContent(relativePath string, content string) {
	tmpfs.ctrl.TestingT().Helper()

	tmpfs.ctrl.AssertThat(relativePath, tmpfs.FileContent(content), "file content")
}

func (tmpfs *TmpFS) AssertNotFileContent(relativePath string, content string) {
	tmpfs.ctrl.TestingT().Helper()

	tmpfs.ctrl.AssertThat(relativePath, tmpfs.ctrl.Not(tmpfs.FileContent(content)), "file content")
}
