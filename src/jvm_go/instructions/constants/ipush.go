package constants

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

//从操作数中获取一个byte型证书，扩展成int型，然后推入栈顶
type BIPUSH struct {
	val int8
}



func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

//short
type SIPUSH struct {
	val int16
}


func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
