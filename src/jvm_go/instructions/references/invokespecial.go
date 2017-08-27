package references

import (
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef()
}

