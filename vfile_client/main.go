package main

import "C"
import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"visionvera/vfile_client/fileTool"
	"visionvera/vfile_client/proto/vfile"
	"visionvera/vfile_client/utils/config"
	"visionvera/vfile_client/utils/logging"
)

var grpcClient vfile.FileServiceClient

func init() {
	//初始化配置文件
	if err := config.Init(""); err != nil {
		panic(err)
	}
	//日志初始化
	logging.InitAllLogger()
	var serviceHost = config.GetString("server.host") + ":" + config.GetString("server.port")
	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	grpcClient = vfile.NewFileServiceClient(conn)
}

//export ServiceInfo
func ServiceInfo() {
	//rsp, err := grpcClient.ServiceInfo(context.TODO(), &vfile.ServiceInfoRequest{})
	//if err != nil {
	//	return
	//}
	//return
}

//export Upload
func Upload(file *C.char) {
	//rsp, err := client.Upload(context.TODO(), &vfile.UploadRequest{
	//	Metadata: nil,
	//	File:     nil,
	//})
	//fmt.Println(rsp)
}

func main() {
	var serviceHost = config.GetString("server.host") + ":" + config.GetString("server.port")
	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fileName := "/home/tan/vfile.MD"
	//fileBytes, err := fileTool.Read(fileName)
	metaData, err := fileTool.ParseFile(fileName, "vfile")
	fmt.Printf("%v\n", metaData)
	//client := vfile.NewFileServiceClient(conn)
	//rsp, err := client.Upload(context.TODO(), &vfile.UploadRequest{
	//	Metadata: &metaData,
	//	File:     &vfile.File{Content: fileBytes},
	//})
	//fmt.Println(rsp)
}
