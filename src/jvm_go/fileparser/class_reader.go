package fileparser

import "encoding/binary"

/**
类解析的帮助类
 */
type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data) //这个方法
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data) //这个方法
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data) //这个方法
	self.data = self.data[8:]
	return val
}

//读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
	length := self.readUint16()
	readData := make([]uint16, length)
	for i := range readData {
		readData[i] = self.readUint16()
	}
	return readData
}

func (self *ClassReader) readBytes(length uint32) []byte {
	//readData := make([]byte, length)
	//for i := range readData {
	//	readData[i] = self.readUint8()
	//}
	//return readData

	readData := self.data[:length]
	self.data = self.data[length:]
	return readData
}
