package main

import (
	"jvm_go/utils"
	"fmt"
)

func main() {
	cmd := utils.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.ClassName == "" {
		utils.PrintUsage()
	} else {
		newJVM(cmd).start()
	}
}
//
//func startJVM(cmd *utils.Cmd) {
//
//	classpath := classpath.LoadClasspath(cmd.JreOption, cmd.CpOption)
//
//	classLoader := heap.NewClassLoader(classpath,cmd.VerboseClassFlag)
//
//	className := strings.Replace(cmd.ClassName, ".", "/", -1)
//
//	mainClass := classLoader.LoadClass(className)
//
//	mainMethod := mainClass.GetMainMethod()
//
//	if mainMethod != nil {
//		interpret(mainMethod,cmd.VerboseInstFlag,cmd.Args)
//	} else {
//		fmt.Printf("Main method not found in class %s\n", cmd.ClassName)
//	}
//	//printClassInfo(classFile)
//	//fmt.Printf("content:%v", content)
//}
//
////获取类文件中的main方法
//func getMainMethod(classFile *fileparser.ClassFile) *fileparser.MemberInfo {
//	for _, method := range classFile.Methods() {
//		if method.Name() == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
//			return method
//		}
//	}
//	return nil
//}
//func printClassInfo(cf *fileparser.ClassFile) {
//	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
//	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
//	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
//	fmt.Printf("this class: %v\n", cf.ClassName())
//	fmt.Printf("super class: %v\n", cf.SuperClassName())
//	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
//	fmt.Printf("fields count: %v\n", len(cf.Fields()))
//	for _, f := range cf.Fields() {
//		fmt.Printf("  %s\n", f.Name())
//	}
//	fmt.Printf("methods count: %v\n", len(cf.Methods()))
//	for _, m := range cf.Methods() {
//		fmt.Printf("  %s\n", m.Name())
//	}
//
//}
