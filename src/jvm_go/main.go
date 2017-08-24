package main

import (
	"jvm_go/utils"
	"fmt"
	"jvm_go/classfile"
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

	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.CpOption, cmd.ClassName, cmd.Args)

	classpath := classfile.LoadClasspath(cmd.JreOption, cmd.CpOption)

	content, _, _ := classpath.ReadClass(cmd.ClassName)

	fmt.Printf("content:%v", content)
}
