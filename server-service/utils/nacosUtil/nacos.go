/*
@Time : 2022/5/27 下午12:51
@Author : tan
@File : nacos
@Software: GoLand
*/

package nacosUtil

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"goService/server-service/utils/config"
)

func ClientConfig() constant.ClientConfig {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.GetString("nacos.registry.nameSpace"),
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/Users/tanxianchen/coding/go/src/goService/log",
		CacheDir:            "/Users/tanxianchen/coding/go/src/goService/cache",
		LogLevel:            "debug",
	}
	return clientConfig
}

func GetInstence(serviceName string) *model.Instance {
	clientConfig := constant.ClientConfig{
		NamespaceId: config.GetString("nacos.nameSpace"),
	}
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: config.GetString("nacos.host"),
			Port:   uint64(config.GetInt("nacos.port")),
		},
	}
	// 创建服务发现客户端
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	params := vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   config.GetString(serviceName),
	}
	instance, err := namingClient.SelectOneHealthyInstance(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	return instance
}
