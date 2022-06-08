/*
@Time : 2022/6/1 上午11:05
@Author : tan
@File : Vfile
@Software: GoLand
*/
package handler

import (
	"context"
	"visionvera/vfile/proto/vfile"
	"visionvera/vfile/utils/constant"
)

type FileHandler struct {
	vfile.UnimplementedFileServiceServer
}

func (f FileHandler) ServiceInfo(ctx context.Context, request *vfile.ServiceInfoRequest) (*vfile.ServiceInfoResponse, error) {
	serviceInfoResponse := &vfile.ServiceInfoResponse{
		Version:        vfile.FileServiceVersion(constant.Version),
		Os:             *constant.OS,
		Hostname:       *constant.HostName,
		AcceptProtocol: *constant.AcceptProtocol,
	}
	return serviceInfoResponse, nil
}

func (f FileHandler) Upload(ctx context.Context, request *vfile.UploadRequest) (*vfile.UploadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) UploadStream(server vfile.FileService_UploadStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) UploadWithOption(server vfile.FileService_UploadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) Download(ctx context.Context, request *vfile.DownloadRequest) (*vfile.DownloadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) DownloadStream(request *vfile.DownloadRequest, server vfile.FileService_DownloadStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) DownloadWithOption(request *vfile.DownloadRequest, server vfile.FileService_DownloadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}
