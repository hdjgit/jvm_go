package stack

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

type POP struct {
	base.NoOperandsInstruction
}

//弹出栈顶元素
/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

//弹出double等占用两个字节的变量
/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}