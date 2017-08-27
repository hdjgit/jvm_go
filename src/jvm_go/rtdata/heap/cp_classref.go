package heap

import "jvm_go/fileparser"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *fileparser.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.ClassName()
	return ref
}
