package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
)

// @title ToDoList API
// @version 2.0
// @description the server is developed by Go
// @name dengzhaowork@gmail.com
// @BasePath /api/v1
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	conf.Init()
	r := routes.NewRoute()
	r.Run(conf.HttpPort)
}
