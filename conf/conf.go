package conf

import (
	"ToDoList/cache"
	"ToDoList/model"
	"ToDoList/pkg/utils"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		utils.LogrusObj.Info("config file err,please check config's path")
		panic(err)
	}
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		utils.LogrusObj.Info("zh-cn.yaml: ", err)
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	dns := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	model.DataBase(dns)
	cache.Redis()
}

func LoadServer(file *ini.File) {
	section := file.Section("service")
	AppMode = section.Key("AppMode").String()
	HttpPort = section.Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	section := file.Section("mysql")
	Db = section.Key("Db").String()
	DbHost = section.Key("DbHost").String()
	DbPort = section.Key("DbPort").String()
	DbUser = section.Key("DbUser").String()
	DbPassWord = section.Key("DbPassWord").String()
	DbName = section.Key("DbName").String()
}
