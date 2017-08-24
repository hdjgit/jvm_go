package fileparser

import "jvm_go/log"

type MarkerAttribute struct {
	name string
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	log.Debugf("read marker attribute:%s", self.name)
}
