package lang

import (
	"math"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/Double").
		Register(doubleToRawLongBits, "(D)J").
		Register(longBitsToDouble, "(J)D")
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.GetDoubleVar(0)

	// todo
	bits := math.Float64bits(value)
	frame.PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *rtda.Frame) {
	bits := frame.GetLongVar(0)

	// todo
	value := math.Float64frombits(uint64(bits))
	frame.PushDouble(value)
}
