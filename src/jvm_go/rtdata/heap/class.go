package heap

import "jvm_go/fileparser"

type Class struct {
	accessFlgs        uint16
	name              string //thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *fileparser.ClassFile) *Class {
	class:=&Class{}
	class.accessFlgs=cf.AccessFlags()
	class.name=cf.ClassName()
	class.superClassName=cf.SuperClassName()
	class.interfaceNames=cf.InterfaceNames()
	class.constantPool=newConstantPool(class,cf.ConstantPool())
	class.fields=newFields(class,cf.Fields())
	class.methods=newMethods(class,cf.Methods())
	return class
}
