package rtdata

import "math"

//局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if (maxLocals > 0) {
		return make([]Slot, maxLocals)
	}
	return nil
}

//操作局部变量表和操作数栈
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

//float 先变成int 在设置
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self.SetInt(index, int32(bits))
}

func (self LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self[index].num))
}

func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)         //低位
	self[index+1].num = int32(val >> 32) //高位
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

//double 先变成 long
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	val := self.GetLong(index)
	return math.Float64frombits(uint64(val))
}

//引用
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
