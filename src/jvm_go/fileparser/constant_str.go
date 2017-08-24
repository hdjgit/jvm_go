package fileparser

//用于表示java.lang.String字面量，它本身并不存放字符串数据，值存了常量值索引
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
	utf8Str     string
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
