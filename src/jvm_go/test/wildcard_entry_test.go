package test

import (
	"testing"
	"jvm_go/classfile"
	"fmt"
)

func TestWildcardEntry(t *testing.T) {
	entry := classfile.NewWildcardEntry("/Users/hdj/gitPro/im-network/target/*")
	content, _, _ := entry.ReadClass("com/mogujie/im/common/service/PacketIdLookupService.class")
	fmt.Printf("classContent:%v", content)
}
