package fileparser

import "fmt"

type ClassFile struct {
	magic        uint32 //魔数
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
将字节流转成ClassFile
 */
func Parse(data []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	classReader := &ClassReader{data}
	classFile := &ClassFile{}
	classFile.read(classReader)
	return
}

//从reader中读取信息
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
}

//读取魔数并检查
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	fmt.Printf("magic number is:%X", self.magic)
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: illegal magic!")
	}
}
