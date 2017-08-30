package heap

import "jvm_go/fileparser"

//字段和方法都属于类成员，他们有一些相同的信息（访问标志、名字、描述符）
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
}

//从class文件中拷贝数据
func (self *ClassMember) copyMemberInfo(memberInfo *fileparser.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}


func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}
func (self *ClassMember) Signature() string {
	return self.signature
}
func (self *ClassMember) AnnotationData() []byte {
	return self.annotationData
}
func (self *ClassMember) AccessFlags() uint16 {
	return self.accessFlags
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
