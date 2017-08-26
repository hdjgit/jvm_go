package math

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

//左移指令

type ISHL struct {
	base.NoOperandsInstruction //int 左位移
}

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//int变量只有32位，所以取v2 的前5个bit位即可
	//go右侧必须是unsign，所以需要进行类型转换
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}