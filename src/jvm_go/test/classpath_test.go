package test

import (
	"testing"
	"jvm_go/classpath"
	"fmt"
)

func TestDirEntry(t *testing.T) {
	dirEntry := classpath.NewDirEntry("/Users/hdj/gitPro/im-network/target/classes")
	classContent, _, err := dirEntry.ReadClass("com/mogujie/im/net/DefaultChannelInitializer.class")
	if err != nil {
		panic(err)
	}
	fmt.Printf("classContent:%v", classContent)
}

func TestExists(t *testing.T) {
	exist := classpath.Exists("/home/sss")
	fmt.Print(exist)
}
