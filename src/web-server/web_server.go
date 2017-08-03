package web_server

import (
	"github.com/gin-gonic/gin"
	"web-server/handlers"
)

func StartServer() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	handlerMap(engine)
	engine.Run(":8080")
}

func handlerMap(engine *gin.Engine) {
	engine.GET("/", handlers.GetCallInfos)
	engine.GET("/interfaces/getRpcInfo/:name", handlers.GetOneRpcCallInfo)
	engine.POST("/interfaces/postRpc/:name", handlers.PostRpcCall)
}