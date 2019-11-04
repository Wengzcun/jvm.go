package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("jdk/internal/reflect/NativeConstructorAccessorImpl").
		Register(newInstance0, "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;")
}

// private static native Object newInstance0(Constructor<?> c, Object[] os)
// throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtda.Frame) {
	constructorObj := frame.GetRefVar(0)
	argArrObj := frame.GetRefVar(1)

	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.Class
	if goClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(goClass)
		return
	}

	obj := goClass.NewObj()
	frame.PushRef(obj)

	// call <init>
	args := convertArgs(obj, argArrObj, goConstructor)
	frame.Thread.InvokeMethodWithShim(goConstructor, args)
}
