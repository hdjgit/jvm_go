package classfile

import (
	"path/filepath"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
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

const PathListSeparator string = string(os.PathListSeparator)

/**
根据不同的路径，返回不同的Entry
 */
func NewEntry(path string) Entry {
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		entry := NewZipEntry(path)
		return entry  //TODO go里的接口类型，似乎要使用指针才可以接受？ 答： 并不是go里的类型需要指针接受，而是实现接口时候，参数用的指针
	}

	if strings.Contains(path, PathListSeparator) {
		return NewCompositeEntry(path)
	}

	return NewDirEntry(path)
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
