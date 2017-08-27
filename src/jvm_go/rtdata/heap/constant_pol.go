package heap

import (
	"jvm_go/fileparser"
	"fmt"
)

/*
运行时常量池主要存放两类信息：字面量（literal）和符号引用 （symbolic reference）
字面量包括整数、浮点数和字符串字面量
符号引用包括类符号引用、字段符号引用、方法符号引用和接口方法符号引用
 */

type Constant interface {
}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

//把class文件中的常量池转成运行常量池
func newConstantPool(class *Class, cfCp fileparser.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *fileparser.ConstantIntegerInfo:
			intInfo := cpInfo.(*fileparser.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *fileparser.ConstantFloatInfo:
			floatInfo := cpInfo.(*fileparser.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *fileparser.ConstantLongInfo:
			longInfo := cpInfo.(*fileparser.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *fileparser.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*fileparser.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *fileparser.ConstantStringInfo:
			stringInfo := cpInfo.(*fileparser.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *fileparser.ConstantClassInfo:
			classInfo := cpInfo.(*fileparser.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *fileparser.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*fileparser.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *fileparser.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*fileparser.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *fileparser.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*fileparser.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}
	}

	return rtCp
}

//根据索引返回常量
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constant at index %d", index))
}
