package main

import (
	"web-server"
	"interfaces"
)

func main() {
	interfaces.RegisterAll()
	web_server.StartServer()
}
