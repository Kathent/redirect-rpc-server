package web_server

import (
	"github.com/gin-gonic/gin"
	"web-server/handlers"
)

func StartServer() {
	engine := gin.Default()

	handlerMap(engine)
	engine.Run(":8080")
}

func handlerMap(engine *gin.Engine) {
	//engine.GET("/interfaces/:rpcName", handlers.GetOneRpcCallInfo)
	engine.POST("/interfaces/postRpc/:name", handlers.PostRpcCall)
}