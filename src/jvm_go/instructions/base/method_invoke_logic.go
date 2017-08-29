package base

import (
	"jvm_go/rtdata"
	"jvm_go/rtdata/heap"
	"fmt"
)

/**
	调用方法:
		invokeFrame: 从哪个帧调用的
 */
func InvokeMethod(invokeFrame *rtdata.Frame, method *heap.Method) {
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	//得到方法参数数目
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			//从调用的栈弹出操作数
			slot := invokeFrame.OperandStack().PopSlot()
			//塞到新的栈里面
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack! TODO
	//if method.IsNative() {
	//	if method.Name() == "registerNatives" {
	//		thread.PopFrame()
	//	} else {
	//		panic(fmt.Sprintf("native method: %v.%v%v\n",
	//			method.Class().Name(), method.Name(), method.Descriptor()))
	//	}
	//}
}
