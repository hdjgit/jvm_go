package heap

import (
	"jvm_go/fileparser"
)

//方法信息相较于field多了字节码，更加复杂一些

type Method struct {
	ClassMember
	maxStack       uint
	maxLocals      uint
	code           []byte
	argSlotCount   uint
	lineNumberTable *fileparser.LineNumberTableAttribute
	exceptionTable ExceptionTable
	parameterAnnotationData []byte                         // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte                         // AnnotationDefault_attribute
	parsedDescriptor        *MethodDescriptor
	exceptions              *fileparser.ExceptionsAttribute // todo: rename
}

func newMethods(class *Class, cfMethods []*fileparser.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {

		methods[i] = newMethod(class, cfMethod)
		//methods[i] = &Method{}
		//methods[i].class = class
		//methods[i].copyMemberInfo(cfMethod)
		//methods[i].copyAttributes(cfMethod)
		//methods[i].calcArgSlotCount()
	}
	return methods
}

func newMethod(class *Class, cfMethod *fileparser.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.parsedDescriptor = md
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++ // `this` reference
	}
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4 // todo
	//本地方法的栈和帧只用于存放参数值
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (self *Method) copyAttributes(cfMethod *fileparser.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			self.class.constantPool)
		self.exceptions = cfMethod.ExceptionsAttribute()
		self.annotationData = cfMethod.RuntimeVisibleAnnotationsAttributeData()
		self.parameterAnnotationData = cfMethod.RuntimeVisibleParameterAnnotationsAttributeData()
		self.annotationDefaultData = cfMethod.AnnotationDefaultAttributeData()
	}
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

//func (self *Method) calcArgSlotCount() {
//	parsedDescriptor := parseMethodDescriptor(self.descriptor)
//	for _, paramType := range parsedDescriptor.parameterTypes {
//		self.argSlotCount++
//		if paramType == "J" || paramType == "D" {
//			self.argSlotCount++
//		}
//	}
//	if !self.IsStatic() {
//		self.argSlotCount++ // `this` reference
//	}
//}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}
func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}

func (self *Method) isConstructor() bool {
	return !self.IsStatic() && self.name == "<init>"
}
func (self *Method) ExceptionTypes() []*Class {
	if self.exceptions == nil {
		return nil
	}

	exIndexTable := self.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := self.class.constantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}

// reflection
func (self *Method) ParameterTypes() []*Class {
	if self.argSlotCount == 0 {
		return nil
	}

	paramTypes := self.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = self.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}
func (self *Method) ReturnType() *Class {
	returnType := self.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return self.class.loader.LoadClass(returnClassName)
}

func (self *Method) ParameterAnnotationData() []byte {
	return self.parameterAnnotationData
}

func (self *Method) isClinit() bool {
	return self.IsStatic() && self.name == "<clinit>"
}
func (self *Method) AnnotationDefaultData() []byte {
	return self.annotationDefaultData
}