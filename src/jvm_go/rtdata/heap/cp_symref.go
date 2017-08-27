package heap

import "jvm_go/fileparser"

//类型符号引用

type SymRef struct {
	cp        *ConstantPool //符号引用所在的运行时常量池指针
	className string        //类的完全限定名
	class     *Class        //解析后的类结构体指针
}


