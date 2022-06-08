package main

/*
#include <stdio.h>
*/
import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"visionvera/vfile/handler"
	"visionvera/vfile/interceptor"
	"visionvera/vfile/proto/vfile"
	"visionvera/vfile/utils/config"
	"visionvera/vfile/utils/constant"
	"visionvera/vfile/utils/db"
	"visionvera/vfile/utils/logging"
)

func init() {
	//初始化配置文件
	if err := config.Init(""); err != nil {
		panic(err)
	}
	//日志初始化
	logging.InitAllLogger()
	// 设置dgraph
	db.DgraphInit()
}

func main() {
	// 获取输入参数
	constant.OS = flag.String("os", "linux", "输入 os")
	constant.HostName = flag.String("host", "127.0.0.1", "输入 host name")
	constant.AcceptProtocol = flag.String("protocol", "vfile", "输入 port")

	listen, err := net.Listen("tcp", config.GetString("server.host")+":"+config.GetString("server.port"))
	if err != nil {
		logging.Logger.Fatal(fmt.Sprintf("failed to listen: %v", err.Error()))
		return
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AccessInterceptor), grpc.StreamInterceptor(interceptor.StreamServerInterceptor))
	vfile.RegisterFileServiceServer(server, new(handler.FileHandler))

	defer func() {
		server.Stop()
		err := listen.Close()
		if err != nil {
			logging.Logger.Fatal(fmt.Sprintf("failed to serve: %v", err.Error()))
			return
		}
	}()
	logging.Logger.Print("Serving ...")
	if err := server.Serve(listen); err != nil {
		logging.Logger.Fatal(fmt.Sprintf("failed to serve: %v", err.Error()))
		return
	}
}
