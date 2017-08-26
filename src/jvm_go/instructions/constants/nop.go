package constants

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

//什么都不做的指令
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtdata.Frame) {
	//do nothing
}
