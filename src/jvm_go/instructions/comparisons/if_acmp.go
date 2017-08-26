package comparisons

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)
//根据引用是否相同进行跳转

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtdata.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtdata.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtdata.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

