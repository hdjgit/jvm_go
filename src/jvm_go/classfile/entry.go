package classfile

import (
	"path/filepath"
	"io/ioutil"
	"fmt"
)

/**
表示类路径
1）提供读取文件内容接口
2）显示类路径
 */
type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

/**
目录类路径
 */
type DirEntry struct {
	dirPath string
}

func NewDirEntry(dirPath string) *DirEntry {
	path, err := filepath.Abs(dirPath)
	if err != nil {
		panic(err)
	}
	dirEntry := &DirEntry{dirPath: path}
	return dirEntry
}

func (self *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	filePath := filepath.Join(self.dirPath, className)
	fileContent, err := ioutil.ReadFile(filePath)
	return fileContent, self, err
}

func (self *DirEntry) String() string {
	return fmt.Sprintf("dir path:%s", self.dirPath)
}
