package fileparser

//类信息，存了一个类名的索引
type ConstantClassInfo struct {
	cp         ConstantPool
	classIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
}

func (self *ConstantClassInfo) ClassName(reader *ClassReader) string {
	return self.cp.getUtf8(self.classIndex)
}
