package main

import (
	"jvm_go/instructions/base"
	"jvm_go/instructions"
	"fmt"
	"jvm_go/rtdata"
)


func interpret(thread *rtdata.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtdata.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtdata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}


//解释器
//func interpret(method *heap.Method, logInst bool, args [] string) {
//	thread := rtdatata.NewThread()
//	frame := thread.NewFrame(method)
//	thread.PushFrame(frame)
//
//	jArgs := createArgsArray(method.Class().Loader(), args)
//	frame.LocalVars().SetRef(0, jArgs)
//
//	defer catchErr(thread) //因为没有实现return指令，所以执行过程必定会出错
//	loop(thread, logInst)
//}
//
//func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
//	stringClass := loader.LoadClass("java/lang/String")
//	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
//	jArgs := argsArr.Refs()
//	for i, arg := range args {
//		jArgs[i] = heap.JString(loader, arg)
//	}
//	return argsArr
//}
//
//func loop(thread *rtdatata.Thread, logInst bool) {
//	reader := &base.BytecodeReader{}
//	for {
//		frame := thread.CurrentFrame()
//		pc := frame.NextPC()
//		thread.SetPC(pc)
//
//		//decode
//		reader.Reset(frame.Method().Code(), pc)
//		opcode := reader.ReadUint8()
//		inst := instructions.NewInstruction(opcode)
//		inst.FetchOperands(reader)
//		frame.SetNextPC(reader.PC())
//
//		//execute
//		if logInst {
//			logInstruction(frame, inst)
//		}
//		inst.Execute(frame)
//		if thread.IsStackEmpty() {
//			break
//		}
//	}
//}
//func logInstruction(frame *rtdatata.Frame, inst base.Instruction) {
//	method := frame.Method()
//	className := method.Class().Name()
//	methodName := method.Name()
//	pc := frame.Thread().PC()
//	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
//}
//func catchErr(thread *rtdatata.Thread) {
//	if r := recover(); r != nil {
//		//fmt.Printf("localVars:%+v\n", frame.LocalVars())
//		//fmt.Printf("OperandStack:%+v\n", frame.OperandStack())
//		logFrames(thread)
//		panic(r)
//	}
//}
//func logFrames(thread *rtdatata.Thread) {
//	for !thread.IsStackEmpty() {
//		frame := thread.PopFrame()
//		method := frame.Method()
//		className := method.Class().Name()
//		fmt.Printf(">> pc:%4d %v.%v%v \n",
//			frame.NextPC(), className, method.Name(), method.Descriptor())
//	}
//}
