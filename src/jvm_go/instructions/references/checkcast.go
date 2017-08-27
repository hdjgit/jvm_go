package references

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
	"jvm_go/rtdata/heap"
)

//check cast 和 instanceOf 区别是，instanceOf会去改变操作数栈
//弹出对象引用，推入判断结果，checkcast 判断失败会直接抛出异常
// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
