package fileparser

//各个属性表达信息不同，和常量池一样不能用统一结构来定义
//使用属性名区分不同的属性
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//读取一组属性
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCounts := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCounts)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

//读某个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp) //构造attrInfo
	attrInfo.readInfo(reader)                           //读取数据进去
	return attrInfo
}
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{MarkerAttribute{name: attrName}}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{MarkerAttribute{name: attrName}}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
