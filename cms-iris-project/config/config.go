package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName   string `json:"app_name"`
	Port      string `json:"port"`
	StaticPth string `json:"static_pth"`
	Mode      string `json:"mode"`
}

var ServerConfig AppConfig

func InitConfig() *AppConfig {
	// 读取配置文件
	file, err := os.Open("../config.json")

	if err != nil {
		panic(err.Error())
	}
	// 创建一个json解码器
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	// 把配置信息转换到结构体
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	// 返回
	return &conf
}
