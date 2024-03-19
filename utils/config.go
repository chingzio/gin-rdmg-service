package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassWord string
	DbName     string

	AppName string
	Port    string

	ConfigPath string
)

func init() {
	file, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadData(file)
	LoadApp(file)
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustInt(13306)
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("gin-rdmg-service")
}

func LoadApp(file *ini.File) {
	AppName = file.Section("app").Key("AppName").MustString("gin-rdmg-service")
	Port = file.Section("app").Key("Port").MustString(":9000")
}
