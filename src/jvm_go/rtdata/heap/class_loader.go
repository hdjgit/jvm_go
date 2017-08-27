package heap

import (
	"jvm_go/classpath"
	"fmt"
	"jvm_go/fileparser"
)

//一个类加载器，只会加载一遍某个类
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class //已经加载的类
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class //类已经加载过了
	}
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

//类的链接,分为验证和准备两个必要阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}
func prepare(class *Class) {
	//todo
}
func verify(class *Class) {
	//todo
}
func (self *ClassLoader) defineClass(bytes []byte) *Class {
	class := parseClass(bytes)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}

}
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		//使用自己的类加载器加载父类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := fileparser.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

//通过classPath加载字节码
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}
