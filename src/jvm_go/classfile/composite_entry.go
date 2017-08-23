package classfile

import (
	"strings"
	"errors"
)

/**
compositeEntry 就是多个entry，使用数组
 */
type CompositeEntry []Entry

func NewCompositeEntry(pathList string) CompositeEntry {

	compositeEntry := []Entry{}

	if !strings.Contains(pathList, PathListSeparator) {
		panic("illegal pathList for CompositeEntry!")
	}
	for _, childPath := range strings.Split(pathList, PathListSeparator) {
		childEntry := NewEntry(childPath)
		compositeEntry = append(compositeEntry, childEntry)
	}
	return compositeEntry
}

func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		content, childEntry, err := entry.ReadClass(className)
		if err != nil {
			return content, childEntry, err
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for _, entry := range self {
		strs = append(strs, entry.String())
	}
	return strings.Join(strs, PathListSeparator)
}
