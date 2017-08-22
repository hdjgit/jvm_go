package utils

import (
	"flag"
	"fmt"
	"os"
)

/*
	定义命令行结构体
	将命令行中的参数映射到这个结构体中
 */

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	ClassName       string
	Args        []string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.ClassName = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
