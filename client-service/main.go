package main

import (
	"github.com/gin-gonic/gin"
	"goService/client-service/router"
	"goService/client-service/utils/config"
	"goService/client-service/utils/db"
	"goService/client-service/utils/logging"
	"goService/client-service/utils/redis"
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
