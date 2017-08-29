package heap

import (
	"jvm_go/classpath"
	"fmt"
	"jvm_go/fileparser"
)

//一个类加载器，只会加载一遍某个类
type ClassLoader struct {
	cp          *classpath.ClassPath
	classMap    map[string]*Class //已经加载的类
	verboseFlag bool
}

func NewClassLoader(cp *classpath.ClassPath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		classMap:    make(map[string]*Class),
		verboseFlag: verboseFlag,
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (self *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		self.loadPrimitiveClass(primitiveType)
	}
}

//加载基本类型
/*
1) void 的基本类型的类名就是void
2）基本类型的类没有超类，也没有实现任何接口
3）非基本类型的类对象是通过ldc指令加载到操作数栈中的
而基本类型的类对象，编译后的指令是getstatic
 */
func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
}

func (self *ClassLoader) loadBasicClasses() {
	//加载Class会加载Object等类
	jlClassClass := self.LoadClass("java/lang/Class")
	//将这些类的extra设置成class
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}


func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class //类已经加载过了
	}
	if name[0] == '[' {
		return self.loadArrayClass(name)
	}
	return self.loadNonArrayClass(name)
}

//数组类直接在内存中生成
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //TODO
		name:        name,
		loader:      self,
		initStarted: true,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

//普通类通过读取文件来加载类
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

//类的链接,分为验证和准备两个必要阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}
func prepare(class *Class) {
	//todo
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//计算实例字段的个数，同时给他们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	//获取父类slot数目
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		//静态和final的就直接初始化
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
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
		panic(fmt.Sprintf("java.lang.ClassFormatError:%+v", err))
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
