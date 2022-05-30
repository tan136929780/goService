/*
@Time : 2022/5/27 上午11:00
@Author : tan
@File : router
@Software: GoLand
*/
package router

import (
	"github.com/gin-gonic/gin"
	"goMicroService/client-service/controller"
)

func UserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/create", controller.CreateUser)
		user.POST("/update", controller.UpdateUser)
		user.POST("/select", controller.SelectUser)
		user.POST("/delete", controller.DeleteUser)
	}
}
