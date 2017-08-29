package lang

import (
	"jvm_go/native"
	"jvm_go/rtdata"
	"math"
)

const jlFloat = "java/lang/Float"

func init() {
	native.Register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtdata.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value) // todo
	frame.OperandStack().PushInt(int32(bits))
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtdata.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack().PushFloat(value)
}

