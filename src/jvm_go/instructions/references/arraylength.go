package references

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	//直接获取数组引用，以获取对应数组的长度
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
