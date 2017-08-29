package heap

import (
	"jvm_go/fileparser"
	"fmt"
)

//方法信息相较于field多了字节码，更加复杂一些

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
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
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		fmt.Printf("native method:%+v",method)
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
