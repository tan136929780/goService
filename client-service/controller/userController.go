/*
@Time : 2022/5/27 上午11:08
@Author : tan
@File : userController
@Software: GoLand
*/
package controller

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/plugins/registry/nacos"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"goService/client-service/proto/proto"
	"goService/client-service/utils/config"
	"net/http"
)

func CreateUser(c *gin.Context) {
	userDto := &proto.UserRequest{}
	err := c.BindJSON(userDto)
	if err != nil {
		fmt.Println(err.Error())
	}
	//instance := nacosCnf.GetInstence("go.micro.service")
	reg := nacos.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{config.GetString("nacos.host") + ":" + config.GetString("nacos.port")}
	})
	microService := micro.NewService(
		micro.Name(config.GetString("client.name")),
		micro.Registry(reg),
		micro.Address(config.GetString("client.host")+":"+config.GetString("client.port")),
	)
	userService := proto.NewUserService("go.micro.service", microService.Client())
	out, err := userService.CreateUser(c, userDto)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, out)
}

func UpdateUser(c *gin.Context) {

}

func SelectUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
