package test

import (
	"testing"
	"jvm_go/classpath"
	"fmt"
)

func TestZipEntry(t *testing.T)  {
	entry := classpath.NewZipEntry("/Users/hdj/gitPro/im-network/target/im-network-1.2.4-GWP-SNAPSHOT.jar")
	content, _, _ := entry.ReadClass("com/mogujie/im/net/protocol/netty/handler/GwpDispatchHandler.class")
	fmt.Printf("classContent:%v", content)
}
