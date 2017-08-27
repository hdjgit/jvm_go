package rtdata

//java虚拟机规范对java虚拟机栈的约束非常宽松，使用链表数据结构来实现java虚拟机

type Stack struct {
	maxSize uint   //栈的容量
	size    uint   //栈的当前大小
	_top    *Frame //栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

//把帧推入栈顶
func (self *Stack) push(frame *Frame) {
	self.size += 1
	if self.size > self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	frame.lower = self._top
	self._top = frame
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	self.size--
	frameToPop := self._top
	self._top = self._top.lower
	frameToPop.lower = nil
	return frameToPop
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}
