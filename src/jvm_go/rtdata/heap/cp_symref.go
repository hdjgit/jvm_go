package heap

//类型符号引用

type SymRef struct {
	cp        *ConstantPool //符号引用所在的运行时常量池指针
	className string        //类的完全限定名
	class     *Class        //解析后的类结构体指针
}

//类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
