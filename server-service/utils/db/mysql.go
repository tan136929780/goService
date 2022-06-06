/*
@Time : 2021/1/15 下午5:02
@Author : tan
@File : db
@Software: GoLand
*/
package db

import (
	"goService/server-service/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MysqlClient *gorm.DB

func MysqlInit() {
	var err error
	MysqlClient, err = gorm.Open(mysql.Open(config.GetString("db.mysql.host")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := MysqlClient.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute)
}
