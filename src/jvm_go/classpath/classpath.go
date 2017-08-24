package classpath

import (
	"os"
	"path/filepath"
	"fmt"
	"strings"
)

/**
	类路径
 */
type ClassPath struct {
	//启动类路径 -XJre 或者 JAVA_HOME/jre
	BootStrapClasspath Entry
	//jre/lib/ext
	ExtensionClasspath Entry
	//-cp 指定
	UserClasspath Entry
}

//加载类路径
func LoadClasspath(jreOption string, cpOption string) *ClassPath {
	classpath := &ClassPath{}
	classpath.loadBootstrapAndExtClasspath(jreOption)
	classpath.loadUserClasspath(cpOption)
	return classpath
}

func (self *ClassPath) loadUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.UserClasspath = NewEntry(cpOption)
}

/**
	1）从jreOption中看有没有
	2）从JAVA_HOME中找到
 */
func (self *ClassPath) loadBootstrapAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	fmt.Printf("jreDir:%s\n", jreDir)

	jreLibPath := filepath.Join(jreDir, "lib", "*")

	fmt.Printf("jreLibPath:%s\n", jreLibPath)
	self.BootStrapClasspath = NewWildcardEntry(jreLibPath)

	extLibPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.ExtensionClasspath = NewWildcardEntry(extLibPath)
}

//找到jre 位置
func getJreDir(jreOption string) string {
	if jreOption != "" && Exists(jreOption) {
		return jreOption
	}

	//2、当前位置
	if (Exists("./jre")) {
		return "./jre"
	}

	//3、环境变量获取
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("can't find jre path!")
}

/*
判断路径是否存在
 */
func Exists(path string) bool {
	//判断一个文件是否存在
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	fileName := strings.Replace(className, ".", "/", -1) + ".class"

	if content, entry, err := self.BootStrapClasspath.ReadClass(fileName); err == nil {
		fmt.Printf("read class from bootstrap!")
		return content, entry, nil
	}

	if content, entry, err := self.ExtensionClasspath.ReadClass(fileName); err == nil {
		fmt.Printf("read class from extension!")
		return content, entry, nil
	}

	fmt.Printf("read class from classpath!")
	return self.UserClasspath.ReadClass(fileName)
}
