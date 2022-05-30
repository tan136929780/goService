package main

import (
	"github.com/gin-gonic/gin"
	"goMicroService/client-service/router"
	"goMicroService/client-service/utils/config"
	"goMicroService/client-service/utils/db"
	"goMicroService/client-service/utils/logging"
	"goMicroService/client-service/utils/redis"
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
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	router.UserRouter(r)
	// 3.监听端口
	r.Run(":" + config.GetString("server.port"))
}
