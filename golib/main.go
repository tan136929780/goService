/*
@Time : 2022/6/2 上午11:46
@Author : tan
@File : main
@Software: GoLand
*/
package main

import "C"
import "lib/hello"

func main() {
}

//export Hello
func Hello(call string) *C.char {
	return C.CString(hello.GetString(call))
}
