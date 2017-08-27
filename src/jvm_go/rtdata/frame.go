package rtdata

import "jvm_go/rtdata/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars     //局部变量表
	operandStack *OperandStack //操作数栈
	thread       *Thread
	nextPC       int
	method       *heap.Method
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) Method() *heap.Method {
	return self.method
}