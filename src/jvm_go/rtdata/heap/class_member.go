package heap

import "jvm_go/fileparser"

//字段和方法都属于类成员，他们有一些相同的信息（访问标志、名字、描述符）
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
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
