package control

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

// switch case： 如果case值可以编码成一个索引表，则实现成tableswitch指令，否则实现成lookupswitch

type TABLE_SWITCH struct {
	defaultOffset int32 //默认情况下执行跳转所需的字节码偏移量
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()  //tableswitch指令操作码的后面有0~3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}


func (self *TABLE_SWITCH) Execute(frame *rtdata.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}