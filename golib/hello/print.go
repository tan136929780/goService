/*
@Time : 2022/6/2 下午1:46
@Author : tan
@File : print
@Software: GoLand
*/
package hello

import "fmt"

type testString struct {
	name string
}

type testBytes struct {
	data []byte
}

func GetString(call string) string {
	a := testString{name: call}
	fmt.Println(a.name)
	return "Hello " + call
}

func GetBytes(call []byte) []byte {
	a := testBytes{data: call}
	fmt.Println(a.data)
	return call
}
