package rpc_client

import (
	"net/rpc"
)

var client *rpc.Client

func GetRpcClient() (*rpc.Client, error){
	if  client == nil{
		c, err := rpc.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			return nil, err
		}

		client = c
		return client, nil
	}

	return client, nil
}
