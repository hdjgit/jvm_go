package heap

import "jvm_go/fileparser"

type Class struct {
	accessFlgs uint16
	name string //thisClassName
	superClassName string
	interfaceNames []string
	constantPool *fileparser.ConstantPool
	fields []*fileparser.Field
}
