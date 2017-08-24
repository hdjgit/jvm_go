package fileparser

type ConstantPool []ConstantInfo

/**
读取ConstantPool
1、表头给出的常量池大小比实际大1，  如果是n，则实际大小是n-1
2、有效索引是1~n-1，0是无效索引
3、Constant_Long_info 和 Constant_Double_info各占两个位置

 */
func ReadConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := int(reader.readUint16())
	cp := make([]ConstantInfo, constantPoolCount)

	for i := 1; i < constantPoolCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
	}
	return cp
}

//获取常量信息
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := self[index]; info != nil {
		return info
	}
	panic("Invalid ConstantPool index")
}

//获取utf8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	info := self.getConstantInfo(index)
	utf8Info, ok := info.(*ConstantUtf8Info)
	if !ok {
		panic("index is not constant utf8")
	}
	return utf8Info.str
}
