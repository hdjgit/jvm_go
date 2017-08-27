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
