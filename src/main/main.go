package main

import (
	"web-server"
	"interfaces"
)

func main() {
	//测试一下git remote
	interfaces.RegisterAll()
	web_server.StartServer()
}
