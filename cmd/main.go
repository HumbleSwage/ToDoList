package main

import (
	"ToDoList/config"
	"ToDoList/pkg/utils"
	"ToDoList/repository/cache"
	"ToDoList/repository/db/dao"
	"ToDoList/router"
)

func main() {
	loading()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
}

func loading() {
	config.InitConfig()
	dao.InitMySQL()
	cache.InitRedis()
	utils.InitLog()
}
