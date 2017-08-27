package heap

import (
	"jvm_go/fileparser"
	"strings"
)

type Class struct {
	accessFlags       uint16
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
	class.accessFlags =cf.AccessFlags()
	class.name=cf.ClassName()
	class.superClassName=cf.SuperClassName()
	class.interfaceNames=cf.InterfaceNames()
	class.constantPool=newConstantPool(class,cf.ConstantPool())
	class.fields=newFields(class,cf.Fields())
	class.methods=newMethods(class,cf.Methods())
	return class
}

// jvms 5.4.4
//判断是否可以访问，就是判断是不是public，然后包名是否一样
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}