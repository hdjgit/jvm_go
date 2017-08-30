package misc

import (
	"jvm_go/native"
	"jvm_go/instructions/base"
	"jvm_go/rtdata"
	"jvm_go/rtdata/heap"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *rtdata.Frame) { // hack: just make VM.savedProps nonempty
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")

	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)

	//classLoader := frame.Method().Class().Loader()
	//jlSysClass := classLoader.LoadClass("java/lang/System")
	//initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	//base.InvokeMethod(frame, initSysClass)
}
