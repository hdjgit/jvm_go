package comparisons

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

/*
ifeq:x==0
ifne:x!=0
iflt:x<0
ifle:x<=0
ifgt:x>0
ifge:x>=0
 */

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
