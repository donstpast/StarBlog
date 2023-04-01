package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 做一些数据处理Ï,初始化配置文件参数
// 引入变量
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	JwtKey     string
)

// 包初始化函数接口
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径下该文件是否存在")
	}
	LoadServer(file)
	LoadDb(file)

}

// LoadServer 采用函数方式抽离Server内参数，初始化server分区内变量内容
func LoadServer(file *ini.File) {
	//ini语法，加载父子分区，采用MustString方法，如果取不到值，则输出为默认值(下面两个相同)
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8383")
	JwtKey = file.Section("server").Key("JwtKey").MustString("ZG9uc3RwYXN0")
}

// LoadDb 采用函数方式抽离Database，初始化database分区内变量内容
func LoadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("star")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("starblog")
	DbName = file.Section("database").Key("DbName").MustString("starblog")

}
