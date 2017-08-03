package handlers

import (
	"github.com/gin-gonic/gin"
	"interfaces"
	"net/http"
	"fmt"
	"web-server/rpc-client"
	"log"
)

func GetOneRpcCallInfo(c *gin.Context){
	log.Printf("enter GetOneRpcCallInfo..%v \n", c.Request)
	name := c.Param("name")
	call := interfaces.GetRpcCall(name)

	c.HTML(http.StatusOK, "get_rpc_call_info.tmpl", gin.H{"title": call.Name,
		"slice": call.ArgFieldName, "name": call.Name})
}

func GetCallInfos(c *gin.Context) {
	log.Printf("enter GetOneRpcCallInfo..%v \n", c.Request)
	callMap := interfaces.GetAllCalls()
	c.HTML(http.StatusOK, "get_calls_info.tmpl", gin.H{"title": "All Calls",
		"map": callMap})
}

func PostRpcCall(c *gin.Context){
	//TODO 解决拷贝的问题
	log.Printf("enter posRpcCall..%v , %v\n", c.Request, c.Request.PostForm)
	name := c.Param("name")
	call := interfaces.GetRpcCall(name)

	if call != nil{
		client, err := rpc_client.GetRpcClient()
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("call err..%s, %v", name, err))
			log.Println(err)
			return
		}

		c.Bind(call.Args)

		log.Println("enter posRpcCal args:", call.Args)
		err = client.Call(call.Name, call.Args, call.Reply)
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("call err..%s, %v", name, err))
			panic(err)
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("cal suc. %s, %v", name, err))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("call not exist..%s", name))
}
