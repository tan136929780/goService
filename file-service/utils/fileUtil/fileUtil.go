/*
@Time : 2022/6/14 下午5:55
@Author : tan
@File : fileRead
@Software: GoLand
*/
package fileUtil

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goService/file-service/proto/newvcms"
	"goService/file-service/proto/vfile"
	"goService/file-service/utils/config"
	"goService/file-service/utils/constant"
	"goService/file-service/utils/errorUtil"
	"goService/file-service/utils/grpcClient"
	"goService/file-service/utils/logging"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type FileInstance struct {
	Uid                  string   `json:"uid"`
	InstanceIdentifier   string   `json:"instance.identifier"`
	InstanceIsa          []string `json:"instance.isa"`
	FileMetaDataType     string   `json:"FileMetaData.type"`
	FileMetaDataHash     string   `json:"FileMetaData.hash"`
	FileMetaDataFileSize int      `json:"FileMetaData.fileSize"`
	FileMetaDataFileName string   `json:"FileMetaData.fileName"`
	FileMetaDataUri      string   `json:"FileMetaData.uri"`
}

type FileInstanceList struct {
	Instances []FileInstance `json:"instances"`
}

func FileRead(fileName string) ([]byte, error) {
	pathName := GetStorePath()
	fileName = pathName + strings.Trim(strings.TrimPrefix(fileName, "vfile://"), "/")
	fileExist := CheckFileExist(fileName)
	if !fileExist {
		logging.DownloadLogger.Info(fmt.Sprintf("FileRead: 下载文件不存在， uri: %s", fileName))
		return nil, errorUtil.NewStringError("下载文件不存在")
	}
	fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
	if err != nil {
		logging.DownloadLogger.Error(fmt.Sprintf("FileRead: %s， uri: %s", err.Error(), fileName))
		return nil, err
	}
	defer fp.Close()
	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		logging.DownloadLogger.Error(fmt.Sprintf("FileRead: %s， uri: %s", err.Error(), fileName))
		return nil, err
	}
	return bytes, err
}

func FileWrite(fileBytes []byte, fileName string) (string, error) {
	pathName := GetStorePath()
	_, err := os.Stat(pathName)
	if os.IsNotExist(err) {
		err := os.MkdirAll(pathName, 0755)
		if err != nil {
			logging.UploadLogger.Error(fmt.Sprintf("FileWrite: %s， uri: %s", err.Error(), fileName))
			return "", err
		}
	}
	isExist := CheckFileExist(fileName)
	if isExist {
		logging.UploadLogger.Info(fmt.Sprintf("FileWrite：文件存在直接返回， uri:%s", fileName))
		return fileName, nil
	}
	fileName = pathName + fileName
	err = ioutil.WriteFile(fileName, fileBytes, 0755)
	if err != nil {
		logging.UploadLogger.Info(fmt.Sprintf("FileWrite：%s， uri:%s", err.Error(), fileName))
		return "", err
	}
	return fileName, nil
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFileMd5(fileName string) (md5Str string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	md5Str = hex.EncodeToString(hashInBytes)
	return
}

func GetBytesMd5(fileBytes []byte) string {
	hash := md5.New()
	hash.Write(fileBytes)
	hashInBytes := hash.Sum(nil)
	md5Str := hex.EncodeToString(hashInBytes)
	return md5Str
}

func ParseFile(fileName, proto string) (vfile.MetaData, error) {
	fileNameWithSuffix := path.Base(fileName)
	fileType := path.Ext(fileNameWithSuffix)
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	md5Str, err := GetFileMd5(fileName)
	if err != nil {
		return vfile.MetaData{}, err
	}
	return vfile.MetaData{
		FileName: proto + "://" + fileNameOnly,
		Uri:      "",
		Type:     fileType,
		Hash:     md5Str,
		FileSize: 0,
	}, nil
}

func GetStorePath() string {
	return strings.TrimSuffix(config.GetString("upload.fileWriteDir"), "/") + "/"
}

func CreateFileInstance() (*newvcms.CreateResult, error) {
	isaProperty, _ := json.Marshal([]string{"property"})
	propreties, _ := json.Marshal([]map[string]string{
		{
			"instance.identifier": "文件名",
			"instance.isa":        string(isaProperty),
			"property.identifier": "fileName",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件唯一标识",
			"instance.isa":        string(isaProperty),
			"property.identifier": "uri",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件类型",
			"instance.isa":        string(isaProperty),
			"property.identifier": "type",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "hash",
			"instance.isa":        string(isaProperty),
			"property.identifier": "hash",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件大小",
			"instance.isa":        string(isaProperty),
			"property.identifier": "fileSize",
			"property.type":       "int",
			"property.status":     "1",
		},
	})
	isa, _ := json.Marshal([]string{"type"})
	createInstanceList := make([]*newvcms.InstanceInfo, 0)
	createInstanceList = append(createInstanceList, &newvcms.InstanceInfo{
		TypeIdentifier: "type",
		Values: map[string]string{
			"instance.identifier": "文件上传模型",
			"instance.isa":        string(isa),
			"type.identifier":     "FileMetaData",
			"type.property":       string(propreties),
			"type.status":         "1",
		},
		Uid:          "",
		RelationInfo: nil,
	})
	request := &newvcms.InstanceCreateRequest{Instances: createInstanceList}
	response, err := grpcClient.AddFileInstence(config.GetString("grpc.instanceService.host"), config.GetInt("grpc.instanceService.port"), request)
	if err != nil {
		logging.MetaDataLogger.Error(err)
		return nil, err
	}
	return response, nil
}

func CreateFileMetaData(metaData *vfile.MetaData) (*newvcms.CreateResult, error) {
	fileInstance, err := FindFileMetaData(metaData.Uri)
	if fileInstance != nil {
		logging.MetaDataLogger.Info(fmt.Sprintf("CreateFileMetaData: 文件模型已存在，uri: %s", metaData.Uri))
		return &newvcms.CreateResult{
			Code:    constant.SUCCESS,
			Message: "文件模型已存在",
			Uid:     []string{fileInstance.Uid},
		}, nil
	}
	if err != nil {
		logging.MetaDataLogger.Error(fmt.Sprintf("CreateFileMetaData: %s， uri: %s", err.Error(), metaData.Uri))
		return nil, err
	}
	isa, _ := json.Marshal([]string{"FileMetaData"})
	createInstanceList := make([]*newvcms.InstanceInfo, 0)
	createInstanceList = append(createInstanceList, &newvcms.InstanceInfo{
		TypeIdentifier: "FileMetaData",
		Values: map[string]string{
			"instance.identifier":   metaData.Uri,
			"instance.isa":          string(isa),
			"FileMetaData.fileName": metaData.FileName,
			"FileMetaData.uri":      metaData.Uri,
			"FileMetaData.type":     metaData.Type,
			"FileMetaData.hash":     metaData.Hash,
			"FileMetaData.fileSize": strconv.FormatInt(metaData.FileSize, 10),
		},
		Uid:          "",
		RelationInfo: nil,
	})
	request := &newvcms.InstanceCreateRequest{Instances: createInstanceList}
	response, err := grpcClient.AddFileInstence(config.GetString("grpc.instanceService.host"), config.GetInt("grpc.instanceService.port"), request)
	if err != nil {
		logging.MetaDataLogger.Error(fmt.Sprintf("CreateFileMetaDat: %s， uri: %s", err.Error(), metaData.Uri))
		return nil, err
	}
	return response, nil
}

func FindFileMetaData(uri string) (*FileInstance, error) {
	propertyExpressionList := make([]*newvcms.PropertyExpression, 0)
	propertyExpressionList = append(propertyExpressionList, &newvcms.PropertyExpression{
		Property: "instance.identifier",
		Operator: "eq",
		Value:    uri,
	}, &newvcms.PropertyExpression{
		Property: "FileMetaData.status",
		Operator: "eq",
		Value:    "1",
	})
	conditionArray := make([]*newvcms.SearchCondition, 0)
	conditionArray = append(conditionArray, &newvcms.SearchCondition{
		TypeIdentifier:     "FileMetaData",
		Properties:         nil,
		Page:               0,
		PageSize:           10,
		Sorters:            nil,
		Expression:         "",
		PropertyExpression: propertyExpressionList,
		IsGroupby:          false,
		GroupbyProperty:    "",
		RelationFilters:    nil,
	})
	request := &newvcms.InstanceFindRequest{
		Condition: conditionArray,
	}
	response, err := grpcClient.FindInstance(config.GetString("grpc.instanceService.host"), config.GetInt("grpc.instanceService.port"), request)
	if err != nil {
		logging.MetaDataLogger.Error(fmt.Sprintf("FindFileMetaData: %s， uri: %s", err.Error(), uri))
		return nil, err
	}
	data := &FileInstanceList{}
	if len(response.Result) > 0 {
		err = json.Unmarshal([]byte(response.Result[0].Data), data)
		if err != nil {
			logging.MetaDataLogger.Error(fmt.Sprintf("FindFileMetaData: %s， uri: %s", err.Error(), uri))
			return nil, err
		}
		if len(data.Instances) > 0 {
			logging.MetaDataLogger.Error(fmt.Sprintf("FindFileMetaData: 文件模型已存在，返回历史录入信息， uri: %s", uri))
			return &data.Instances[0], nil
		}
	}
	return nil, nil
}
