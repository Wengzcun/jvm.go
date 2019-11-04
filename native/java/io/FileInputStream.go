package io

import (
	"io"
	"os"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/io/FileInputStream").
		Register(available, "()I").
		Register(close0, "()V").
		Register(readBytes, "([BII)I").
		Register(open0, "(Ljava/lang/String;)V")
}

// public native int available() throws IOException;
// ()I
func available(frame *rtda.Frame) {
	// todo
	frame.PushInt(1)
}

// private native void close0() throws IOException;
// ()V
func close0(frame *rtda.Frame) {
	this := frame.GetThis()

	goFile := this.Extra.(*os.File)
	err := goFile.Close()
	if err != nil {
		// todo
		panic("IOException")
	}
}

// private native void open(String name) throws FileNotFoundException;
// (Ljava/lang/String;)V
func open0(frame *rtda.Frame) {
	this := frame.GetThis()
	name := frame.GetRefVar(1)

	goName := name.JSToGoStr()
	goFile, err := os.Open(goName)
	if err != nil {
		frame.Thread.ThrowFileNotFoundException(goName)
		return
	}

	this.Extra = goFile
}

// private native int readBytes(byte b[], int off, int len) throws IOException;
// ([BII)I
func readBytes(frame *rtda.Frame) {
	this := frame.GetThis()
	buf := frame.GetRefVar(1)
	off := frame.GetIntVar(2)
	_len := frame.GetIntVar(3)

	goFile := this.Extra.(*os.File)
	goBuf := buf.GetGoBytes()
	goBuf = goBuf[off : off+_len]

	// func (f *File) Read(b []byte) (n int, err error)
	n, err := goFile.Read(goBuf)
	if err == nil || n > 0 || err == io.EOF {
		frame.PushInt(int32(n))
	} else {
		// todo
		panic("IOException!" + err.Error())
	}
}
