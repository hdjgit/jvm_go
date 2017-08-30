package reserved

import (
	"jvm_go/instructions/base"
	"jvm_go/native"
	"jvm_go/rtdata"
	_ "jvm_go/native/java/lang"
	_ "jvm_go/native/java/io"
	_ "jvm_go/native/java/security"
	_ "jvm_go/native/sun/misc"
	_ "jvm_go/native/sun/reflect"
	_ "jvm_go/native/java/util/concurrent/atomic"
)

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

//VM.savedProps.setProperty("foo","bar")
func (self *INVOKE_NATIVE) Execute(frame *rtdata.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
