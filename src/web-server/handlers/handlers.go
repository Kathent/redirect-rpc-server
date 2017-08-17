package handlers

import (
	"github.com/gin-gonic/gin"
	"interfaces"
	"net/http"
	"fmt"
	"log"
)

func GetOneRpcCallInfo(c *gin.Context){
	log.Printf("enter GetOneRpcCallInfo..%v \n", c.Request)
	name := c.Param("name")
	call := interfaces.GetRpcCall(name)

	base := call.GetBaseRpcCall()
	c.HTML(http.StatusOK, "get_rpc_call_info.tmpl", gin.H{"title": base.Name,
		"slice": base.ArgFieldName, "name": base.Name})
}

func GetCallInfos(c *gin.Context) {
	log.Printf("enter GetOneRpcCallInfo..%v \n", c.Request)
	callMap := interfaces.GetAllCalls()
	c.HTML(http.StatusOK, "get_calls_info.tmpl", gin.H{"title": "All Calls",
		"map": callMap})
}

func PostRpcCall(c *gin.Context){
	log.Printf("enter posRpcCall..%v , %v\n", c.Request, c.Request.PostForm)
	name := c.Param("name")
	call := interfaces.GetRpcCall(name)

	if call != nil{
		base := call.GetBaseRpcCall()

		args := base.ArgsGen()
		reply := base.ReplyGen()
		c.Bind(args)

		err := call.InvokeRpcCall(args, reply)
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("internal err %v", err))
		}

		c.String(http.StatusOK, fmt.Sprintf("cal return %v", reply))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("call not exist..%s", name))
}
