package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/plugins/registry/nacos"
	registry2 "goService/server-service/registry"
	"goService/server-service/utils/config"
	"goService/server-service/utils/db"
	"goService/server-service/utils/logging"
	"goService/server-service/utils/nacosUtil"
	"goService/server-service/utils/redis"
)

func init() {
	//初始化配置文件
	if err := config.Init(""); err != nil {
		panic(err)
	}
	//日志初始化
	logging.InitAllLogger()
	// 设置mysql
	db.MysqlInit()
	// 设置dgraph
	db.DgraphInit()
	//设置redis
	redis.Setup()

}

func main() {
	//nacos作为服务注册中心
	reg := nacos.NewRegistry(nacos.WithAddress([]string{config.GetString("nacos.host") + ":" + config.GetString("nacos.port")}), nacos.WithClientConfig(nacosUtil.ClientConfig()))

	// 生成服务
	microService := micro.NewService(
		micro.Name(config.GetString("server.name")),
		micro.Registry(reg),
		micro.Address(config.GetString("server.host")+":"+config.GetString("server.port")),
	)
	microService.Init()
	// 注册逻辑处理handler
	registry2.RegistryService(microService.Server())
	err := microService.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
