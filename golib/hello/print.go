/*
@Time : 2022/6/2 下午1:46
@Author : tan
@File : print
@Software: GoLand
*/
package hello

import "fmt"

type test struct {
	name string
}

func GetString(call string) string {
	a := test{name: call}
	fmt.Println(a.name)
	return "Hello " + call
}
