package classfile

/**
	类路径
 */
type ClassPath struct {
	//启动类路径 -XJre 或者 JAVA_HOME/jre
	BootStrapClasspath Entry
	//jre/lib/ext
	ExtensionClasspath Entry
	//-cp 指定
	UserClasspath      Entry
}
