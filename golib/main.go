/*
@Time : 2022/6/2 上午11:46
@Author : tan
@File : main
@Software: GoLand
*/
package main

import "C"
import (
	"goService/golib/hello"
	"unsafe"
)

func main() {
}

//export Hello
func Hello(call string) *C.char {
	return C.CString(hello.GetString(call))
}

//export Download
func Download(raw *C.char, size C.int) *C.char {
	data := C.GoBytes(unsafe.Pointer(raw), size)
	return C.CString(hello.GetBytes(data))
}
