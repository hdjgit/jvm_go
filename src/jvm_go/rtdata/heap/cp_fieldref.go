package heap

import "jvm_go/fileparser"

type FieldRef struct {
	MemberRef
	field *Field
}

//字段解析
func newFieldRef(cp *ConstantPool, refInfo *fileparser.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// jvms 5.4.3.2
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class //常量池所属类
	c := self.ResolvedClass() // 加载了的类
	field := lookupField(c, self.name, self.descriptor) //从类中找到这个field

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

//找自己的字段 -> 找父接口的字段 -> 找父类的字段
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}