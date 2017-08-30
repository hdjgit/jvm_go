package misc

import (
	"jvm_go/native"
	"jvm_go/rtdata"
)

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtdata.Frame) {
	frame.OperandStack().PushRef(nil)
}
