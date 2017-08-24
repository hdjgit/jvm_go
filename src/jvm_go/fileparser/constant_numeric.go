package fileparser

import "math"

//int常量，有个4字节的int数据
//比int更小的boolean、byte、short、char类型的常量都放在CONSTANT_Integer_info中
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val uint64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = bytes
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
