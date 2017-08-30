package main

import (
	"jvm_go/classpath"
	"jvm_go/instructions/base"
	"strings"
	"fmt"
	. "jvm_go/utils"
	"jvm_go/rtdata/heap"
	"jvm_go/rtdata"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtdata.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.JreOption, cmd.CpOption)
	classLoader := heap.NewClassLoader(cp, cmd.VerboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtdata.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.VerboseInstFlag)
}

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.ClassName, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.ClassName)
		return
	}

	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.VerboseInstFlag)
}

func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.Args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.Args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}