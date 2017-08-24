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

	fileparser.Parse(content)

	//fmt.Printf("content:%v", content)
}
