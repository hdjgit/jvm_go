package test

import (
	"testing"
	"jvm_go/classfile"
	"fmt"
)

func TestDirEntry(t *testing.T) {
	dirEntry := classfile.NewDirEntry("/Users/hdj/gitPro/im-network/target/classes")
	classContent, _, err := dirEntry.ReadClass("com/mogujie/im/net/DefaultChannelInitializer.class")
	if err != nil {
		panic(err)
	}
	fmt.Printf("classContent:%v", classContent)
}
