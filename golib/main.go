/*
@Time : 2022/6/2 上午11:46
@Author : tan
@File : main
@Software: GoLand
*/
package main

import "C"
import (
	"bytes"
	"goService/golib/hello"
	"unsafe"
)

func main() {
}

//export Hello1
func Hello1(call string) *C.char {
	return C.CString(hello.GetString(call))
}

//export Hello2
func Hello2(call string) (*C.char, *C.char) {
	return C.CString(hello.GetString(call)), C.CString(hello.GetString(call))
}

//export Hello3
func Hello3(call *C.char, size C.int) (unsafe.Pointer, int) {
	data := C.GoBytes(unsafe.Pointer(call), size)
	return C.CBytes(bytes.Title(data)), len(data)
}
