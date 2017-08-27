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
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
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

func (self ConstantPool) getClassName(index uint16) string {
	info := self.getConstantInfo(index)
	classInfo, ok := info.(*ConstantClassInfo)
	if !ok {
		panic("index is not class name")
	}
	return classInfo.ClassName()
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	info := self.getConstantInfo(index)
	nameAndTypeInfo, ok := info.(*ConstantNameAndType)
	if !ok {
		panic("index is not constant utf8")
	}
	name := self.getUtf8(nameAndTypeInfo.nameIndex)
	_type := self.getUtf8(nameAndTypeInfo.descriptorIndex)
	return name, _type
}
