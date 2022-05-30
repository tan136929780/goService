/*
@Time : 2022/5/27 下午3:53
@Author : tan
@File : userDto
@Software: GoLand
*/
package dto

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
