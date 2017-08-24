package fileparser

//int size
//void run
type ConstantNameAndType struct {
	nameIndex       uint16 //字段或方法名称
	descriptorIndex uint16 //类型描述符 byte、short、char、int  对应 B、S、C、I
}

func (self *ConstantNameAndType) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
