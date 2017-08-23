package classfile

import (
	"path/filepath"
	"os"
	"strings"
)

/*
通配符类加载
 */
//type WildcardEntry struct {
//	path string
//}

func NewWildcardEntry(path string) CompositeEntry {

	dirPath := path[:len(path)-1]

	entry := CompositeEntry{}

	walkFn := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if info.IsDir() && path != dirPath {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
			jarEntry := NewZipEntry(path)
			entry = append(entry, jarEntry)
		}

		return nil
	}

	filepath.Walk(dirPath, walkFn)

	return entry
}
