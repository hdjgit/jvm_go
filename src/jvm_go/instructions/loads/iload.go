package loads

import (
	"jvm_go/rtdata"
	"jvm_go/instructions/base"
)

//从局部变量表中获取一个int变量，然后推入操作数栈栈顶

func _iload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtdata.Frame) {
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *rtdata.Frame) {
	_iload(frame, 0)
}


type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtdata.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtdata.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtdata.Frame) {
	_iload(frame, 3)
}

