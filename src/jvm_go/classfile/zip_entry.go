package classfile

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"fmt"
	"errors"
)

type ZipEntry struct {
	ZipFilePath string //压缩文件的路径
}

func NewZipEntry(zipFilePath string) *ZipEntry {
	absPath, err := filepath.Abs(zipFilePath)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	//1、打开zip文件
	reader, err := zip.OpenReader(self.ZipFilePath)
	if err != nil {
		fmt.Printf("ZipEntry ReadClass error:%v", err)
		return nil, nil, err
	}

	defer reader.Close()

	//2、遍历zip文件，找到className是否有一样的
	for _, f := range reader.File {
		if f.Name == className { //比如 com/mogujie/im/net/protocol/netty/handler/GwpDispatchHandler.class
			fc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer fc.Close()
			data, err := ioutil.ReadAll(fc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	//到了这里说明没有找到这个文件
	return nil, nil, errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string {
	return fmt.Sprintf("zip path:%s", self.ZipFilePath)
}
