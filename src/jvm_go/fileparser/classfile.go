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
	self.readAndCheckVersion(reader)
	self.constantPool = ReadConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

//读取并检查版本好
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	fmt.Printf("majorVersion: %d, minorVersion: %d\n", self.majorVersion, self.minorVersion)
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}

//读取魔数并检查
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	fmt.Printf("magic number is:%X\n", self.magic)
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: illegal magic!")
	}
}
