package native

import "jvm_go/rtdata"

//此虚拟机的jni是用go实现的，就是将java方法转成go方法调用
type NativeMethod func(frame *rtdata.Frame)

var registry = map[string]NativeMethod{}

//建立java方法和go方法的映射
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}


func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	//java.lang.Object通过registerNatives的本地方法来注册其他本地方法的，
	//本书实现，所有注册都是自己实现的，所以这里使用一个空的方法
	if methodDescriptor == "()V" {
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}
	return nil
}

func emptyNativeMethod(frame *rtdata.Frame) {
	// do nothing
}
