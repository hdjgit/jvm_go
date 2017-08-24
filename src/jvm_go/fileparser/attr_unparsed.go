package fileparser

import "jvm_go/log"

//未解析的属性
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	log.Debugf("parse unsupport attr! name:%s", self.name)
	self.info = reader.readBytes(self.length)
}
