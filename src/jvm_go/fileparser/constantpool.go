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
