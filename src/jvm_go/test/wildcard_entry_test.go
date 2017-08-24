package test

import (
	"testing"
	"jvm_go/classpath"
	"fmt"
)

func TestWildcardEntry(t *testing.T) {
	entry := classpath.NewWildcardEntry("/Users/hdj/gitPro/im-network/target/*")
	content, _, _ := entry.ReadClass("com/mogujie/im/common/service/PacketIdLookupService.class")
	fmt.Printf("classContent:%v", content)
}

func TestJre(t *testing.T) {
	entry := classpath.NewWildcardEntry("/Library/Java/JavaVirtualMachines/jdk1.8.0_91.jdk/Contents/Home/jre/lib/*")
	fmt.Printf("length: %d\n",len(entry))
	content, _, _ := entry.ReadClass("java/lang/Object.class")
	fmt.Printf("classContent:%v", content)
}
