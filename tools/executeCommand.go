// Package tools -----------------------------
// @file      : executeCommand.go
// @author    : Penguin
// @contact   : 1920732889@qq.com
// @time      : 2024/6/6 15:00
// -------------------------------------------
package tools

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"pentest_mac_tools/model"
	"strings"
)

// todo 添加一个args参数，用来 -Xdock:icon=Behinder.icns  加这种参数，这版本展示不加了
func ExecuteCommand(yamldata []byte, name, path, optional, command, value string) error {

	var config model.Config
	yaml.Unmarshal(yamldata, &config)
	currentDir, err := os.Getwd()
	if err != nil {
		WriteErrorToFile(err.Error())
		return err
	}
	java8_path := currentDir + config.Java_path.Java8
	java11_path := currentDir + config.Java_path.Java11
	fmt.Println(java8_path)
	var execCommand string
	//处理路径的问题
	if command == "Java8" {
		parts := strings.Split(path, "/")
		startname := parts[len(parts)-1]
		allButLast := parts[:len(parts)-1]
		fullPath := strings.Join(allButLast, "/")
		execCommand = fmt.Sprintf("cd %s &&  %s  %s   %s   %s", fullPath, java8_path, optional, value, startname)
		fmt.Println(execCommand)

	}
	if command == "Java11" {
		parts := strings.Split(path, "/")
		startname := parts[len(parts)-1]
		allButLast := parts[:len(parts)-1]
		fullPath := strings.Join(allButLast, "/")
		execCommand = fmt.Sprintf("cd %s &&  %s  %s  %s   %s", fullPath, java11_path, optional, value, startname)
		fmt.Println(execCommand)
	}
	if command == "Open" {
		parts := strings.Split(path, "/")
		startname := parts[len(parts)-1]
		allButLast := parts[:len(parts)-1]
		fullPath := strings.Join(allButLast, "/")
		execCommand = fmt.Sprintf("cd %s && open %s", fullPath, startname)
	}
	cmd := exec.Command("sh", "-c", execCommand)
	cmd.Start()
	return nil
}
