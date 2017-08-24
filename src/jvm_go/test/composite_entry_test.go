package test

import (
	"testing"
	"jvm_go/classpath"
	"fmt"
)

func TestCompositeEntry(t *testing.T) {
	entry := classpath.NewCompositeEntry("/Users/hdj/gitPro/im-network/target/classes:/Users/hdj/gitPro/im-network/target/im-network-1.3.0-RELEASE.jar")
	content, _, _ := entry.ReadClass("com/mogujie/im/common/service/PacketIdLookupService.class")
	fmt.Printf("classContent:%v", content)
}
