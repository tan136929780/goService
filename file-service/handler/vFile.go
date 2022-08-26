package handler

import (
	"fmt"
	"goService/file-service/proto/vfile"
	"goService/file-service/utils/constant"
	"goService/file-service/utils/fileUtil"
	"goService/file-service/utils/logging"
	"os"
	"path"
	"strings"
)

type VFile struct {
}

func (f VFile) Upload(uploadRequest *vfile.UploadRequest) (*vfile.UploadResponse, error) {
	dataHash := fileUtil.GetBytesMd5(uploadRequest.File.Content)
	if dataHash != uploadRequest.Metadata.Hash {
		return &vfile.UploadResponse{
			Code:    constant.FAILED,
			Message: "文件校验失败",
			Uri:     "",
		}, nil
	}
	fileType := uploadRequest.Metadata.FileName[strings.Index(uploadRequest.Metadata.FileName, ".")+1 : len(uploadRequest.Metadata.FileName)]
	fileName, err := fileUtil.FileWrite(uploadRequest.File.Content, uploadRequest.Metadata.Hash+"."+fileType)
	if err != nil {
		logging.DownloadLogger.Error(fmt.Sprintf("Download: %s，uri: %s", err.Error(), uploadRequest.Metadata.Uri))
		return &vfile.UploadResponse{
			Code:    constant.FAILED,
			Message: "文件上传失败",
			Uri:     "",
		}, nil
	}
	uploadRequest.Metadata.Uri = uploadRequest.Metadata.Hash + "." + fileType
	_, err = fileUtil.CreateFileMetaData(uploadRequest.Metadata)
	if err != nil {
		logging.DownloadLogger.Error(fmt.Sprintf("Download: %s，uri: %s", "创建文件元信息失败删除文件", uploadRequest.Metadata.Uri))
		os.Remove(fileUtil.GetStorePath() + fileName)
		return &vfile.UploadResponse{
			Code:    constant.FAILED,
			Message: "文件信息保存失败",
			Uri:     path.Base(fileName),
		}, nil
	}
	return &vfile.UploadResponse{
		Code:    constant.SUCCESS,
		Message: "上传成功",
		Uri:     path.Base(fileName),
	}, nil
}

func (f VFile) UploadStream(metaData *vfile.MetaData, file *vfile.File) error {
	// TODO implement me
	panic("implement me")
}

func (f VFile) UploadWithOption(server vfile.FileService_UploadWithOptionServer) error {
	// TODO implement me
	panic("implement me")
}

func (f VFile) Download(DownloadRequest *vfile.DownloadRequest) (*vfile.DownloadResponse, error) {
	fileInstance, err := fileUtil.FindFileMetaData(strings.TrimPrefix(DownloadRequest.Uri, "vfile://"))
	if fileInstance == nil {
		logging.DownloadLogger.Info(fmt.Sprintf("Download: 文件不存在，uri: %s", DownloadRequest.Uri))
		return &vfile.DownloadResponse{
			Code:     constant.SUCCESS,
			Message:  "文件不存在",
			Metadata: &vfile.MetaData{},
			File:     &vfile.File{},
		}, nil
	}
	fileBytes, err := fileUtil.FileRead(DownloadRequest.Uri)
	file := &vfile.File{Content: fileBytes}
	if err != nil {
		logging.DownloadLogger.Error(fmt.Sprintf("Download: %s，uri: %s", err.Error(), DownloadRequest.Uri))
		return &vfile.DownloadResponse{
			Code:     constant.FAILED,
			Message:  err.Error(),
			Metadata: &vfile.MetaData{},
			File:     &vfile.File{},
		}, nil
	}
	metaData := &vfile.MetaData{
		FileName: fileInstance.FileMetaDataFileName,
		Uri:      fileInstance.FileMetaDataUri,
		Type:     fileInstance.FileMetaDataType,
		Hash:     fileInstance.FileMetaDataHash,
		FileSize: int64(fileInstance.FileMetaDataFileSize),
	}
	res := &vfile.DownloadResponse{
		Code:     constant.SUCCESS,
		Message:  "下载成功",
		Metadata: metaData,
		File:     file,
	}
	return res, nil
}

func (f VFile) DownloadStream(request *vfile.DownloadRequest, server vfile.FileService_DownloadStreamServer) error {
	// TODO implement me
	panic("implement me")
}

func (f VFile) DownloadWithOption(request *vfile.DownloadRequest, server vfile.FileService_DownloadWithOptionServer) error {
	// TODO implement me
	panic("implement me")
}
