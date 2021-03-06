package atomic

import (
	"jvm_go/native"
	"jvm_go/rtdata"
)

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtdata.Frame) {
	frame.OperandStack().PushBoolean(false)
}
