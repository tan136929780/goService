/*
@Time : 2022/6/1 上午11:05
@Author : tan
@File : Vfile
@Software: GoLand
*/
package handler

import (
	"context"
	"visionvera/vfile_client/proto/vfile"
)

type FileHandler struct{}

func (h *FileHandler) ServiceInfo(ctx context.Context, in *vfile.ServiceInfoRequest, out *vfile.ServiceInfoResponse) error {
	return nil
}

func (h *FileHandler) Upload(ctx context.Context, in *vfile.UploadRequest, out *vfile.UploadResponse) error {
	return nil
}

func (h *FileHandler) UploadStream(ctx context.Context, stream vfile.FileService_UploadStreamStream) error {
	return nil
}

func (h *FileHandler) UploadWithOption(ctx context.Context, stream vfile.FileService_UploadWithOptionStream) error {
	if _, err := stream.Recv(); err != nil {
		return err
	}
	return nil
}

func (h *FileHandler) Download(ctx context.Context, in *vfile.DownloadRequest, out *vfile.DownloadResponse) error {
	return nil
}

func (h *FileHandler) DownloadStream(ctx context.Context, in *vfile.DownloadRequest, stream vfile.FileService_DownloadStreamStream) error {
	m := new(vfile.DownloadStreamResponse)
	if err := stream.Send(m); err != nil {
		return err
	}
	return nil
}

func (h *FileHandler) DownloadWithOption(ctx context.Context, in *vfile.DownloadRequest, stream vfile.FileService_DownloadWithOptionStream) error {
	m := new(vfile.DownloadStreamResponse)
	if err := stream.Send(m); err != nil {
		return err
	}
	return nil
}
