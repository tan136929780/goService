/*
@Time : 2022/6/7 下午5:55
@Author : tan
@File : fileRead
@Software: GoLand
*/
package fileTool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Read(fileName string) ([]byte, error) {
	fileExist := CheckFileExist(fileName)
	if fileExist {
		return nil, nil
	}
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	//把file读取到缓冲区中
	defer file.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		//从file读取到buf中
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return nil, err
		}
		//说明读取结束
		if err == io.EOF {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	return chunk, nil
}

func CheckFileExist(fileName string) bool {
	_, error := os.Stat(fileName)
	if !os.IsNotExist(error) {
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
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}
