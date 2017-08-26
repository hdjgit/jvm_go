package control

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

//goto 进行无条件跳转
// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}

