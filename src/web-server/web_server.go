package web_server

import (
	"github.com/gin-gonic/gin"
	"web-server/handlers"
	"config"
)

func StartServer() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	handlerMap(engine)
	engine.Run(config.GetConfig().Server.Addr)
}

func handlerMap(engine *gin.Engine) {
	engine.GET("/interfaces", handlers.GetCallInfos)
	engine.GET("/interfaces/getRpcInfo/:name", handlers.GetOneRpcCallInfo)
	engine.POST("/interfaces/postRpc/:name", handlers.PostRpcCall)
}