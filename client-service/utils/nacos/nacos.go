/*
@Time : 2022/5/27 下午12:51
@Author : tan
@File : nacos
@Software: GoLand
*/

package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"goService/client-service/utils/config"
)

func NamingClient() naming_client.INamingClient {
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
	return namingClient
}

func GetInstence(serviceName string) *model.Instance {
	clientConfig := constant.ClientConfig{
		NamespaceId: config.GetString(""),
	}
	//serverConfig := []constant.ServerConfig{
	//	{
	//		IpAddr: config.GetString("nacos.host"),
	//		Port:   uint64(config.GetInt("nacos.port")),
	//	},
	//}
	// 创建服务发现客户端
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"clientConfig": clientConfig,
		//"serverConfigs": serverConfig,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	params := vo.SelectOneHealthInstanceParam{
		ServiceName: config.GetString(serviceName),
	}
	instance, err := namingClient.SelectOneHealthyInstance(params)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(11111)
	}
	return instance
}
