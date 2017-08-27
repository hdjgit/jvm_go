package main

import (
	"jvm_go/rtdata"
	"jvm_go/instructions/base"
	"jvm_go/instructions"
	"fmt"
	"jvm_go/rtdata/heap"
)

//解释器
func interpret(method *heap.Method) {
	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame) //因为没有实现return指令，所以执行过程必定会出错
	loop(thread, method.Code())
}
func loop(thread *rtdata.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	fmt.Printf("bytecode：%+X\n", bytecode)
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc:%2d inst:%T %+v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("localVars:%+v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%+v\n", frame.OperandStack())
		panic(r)
	}
}
