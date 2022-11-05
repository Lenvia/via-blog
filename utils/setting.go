package utils

import (
	"fmt"
	ini "gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone        int
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {

	file, err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("配置文件读取错误:", err)
	}
	LoadServer(file)  // 服务器配置
	LoadData(file)  // 数据库配置
	LoadQiniu(file)  // 云空间配置

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8008")
	JwtKey = file.Section("server").Key("JwtKey").MustString("qwerty")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("y724290941")
	DbName = file.Section("database").Key("DbName").MustString("viablog")
}

func LoadQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
