// Package model -----------------------------
// @file      : model.go
// @author    : Penguin
// @contact   : 1920732889@qq.com
// @time      : 2024/6/6 12:29

package model

type Config struct {
	Java_path struct {
		Java8  string `yaml:"Java8"`
		Java11 string `yaml:"Java11"`
		Open   string `yaml:"Open"`
	} `yaml:"javapath"`
}

type Tool struct {
	Path    string `yaml:"path"`
	Command string `yaml:"command"`
}

type Categories struct {
	Category []struct {
		Name string `yaml:"CategoryName"`
		Task []struct {
			Name     string `yaml:"TaskName"`
			PATH     string `yaml:"PATH"`
			VALUE    string `yaml:"VALUE"`
			COMMAND  string `yaml:"COMMAND"`
			Optional string `yaml:"Optional"`
		} `yaml:"Tasks"`
	} `yaml:"Categories"`
}
