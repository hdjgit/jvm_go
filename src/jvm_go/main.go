package main

import (
	"jvm_go/utils"
	"fmt"
	"jvm_go/classpath"
	"jvm_go/fileparser"
)

func main() {
	cmd := utils.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.ClassName == "" {
		utils.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *utils.Cmd) {

	classpath := classpath.LoadClasspath(cmd.JreOption, cmd.CpOption)

	content, _, _ := classpath.ReadClass(cmd.ClassName)

	classFile, err := fileparser.Parse(content)

	if err!=nil{
		panic(err)
	}
	printClassInfo(classFile)
	//fmt.Printf("content:%v", content)
}
func printClassInfo(cf *fileparser.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}

}
