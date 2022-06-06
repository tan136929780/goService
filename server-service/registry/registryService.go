/*
@Time : 2022/5/27 下午7:00
@Author : tan
@File : registryService
@Software: GoLand
*/
package registry

import (
	"fmt"
	"github.com/asim/go-micro/v3/server"
	"goService/server-service/grpc/user"
	"goService/server-service/proto/proto"
)

func RegistryService(server server.Server) {
	// user service
	err := proto.RegisterUserServiceHandler(server, user.NewUserGrpcService())
	if err != nil {
		fmt.Println(err.Error())
	}
}
