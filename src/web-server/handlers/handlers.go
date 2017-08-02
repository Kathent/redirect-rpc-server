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

}

func PostRpcCall(c *gin.Context){
	log.Printf("enter posRpcCall..%v", c.Request)
	name := c.Param("name")
	call := interfaces.GetRpcCall(name)

	log.Println(name, call)
	if call != nil{
		client, err := rpc_client.GetRpcClient()
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("call err..%s, %v", name, err))
			fmt.Println(err)
			return
		}

		c.Bind(call.Args)

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
