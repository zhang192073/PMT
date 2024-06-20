// Package tools -----------------------------
// @file      : tools.go
// @author    : Penguin
// @contact   : 1920732889@qq.com
// @time      : 2024/6/6 13:02
// -------------------------------------------
package tools

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadYAMLFile() ([]byte, error) {
	currentDir, err := os.Getwd()
	//if err != nil {
	//	WriteErrorToFile(err.Error())
	//	return nil, fmt.Errorf("获取当前工作目录出错: %v", err)
	//}
	//filename := filepath.Join(currentDir, "/pentest_mac_tools.app/Contents/Resources/resources/config/config.yaml")
	filename := "resources/config/config.yaml"
	data, err := ioutil.ReadFile(filename)
	WriteErrorToFile(filename)
	WriteErrorToFile(currentDir)
	if err != nil {
		WriteErrorToFile(err.Error())
		return nil, fmt.Errorf("读取文件出错: %v", err)
	}
	return data, nil
}

func WriteErrorToFile(errMsg string) {
	file, err := os.OpenFile("/tmp/errorsss.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开错误日志文件失败:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(errMsg + "\n")
	if err != nil {
		fmt.Println("写入错误日志文件失败:", err)
		return
	}
}
