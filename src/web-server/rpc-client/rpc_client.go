package rpc_client

import (
	"net/rpc"
)

var client *rpc.Client

const (
	MAX_RETRY_TIME = 3
)

func GetRpcClient() (*rpc.Client, error){
	if  client == nil {
		c, err := rpc.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			return nil, err
		}

		client = c
		return client, nil
	}

	return client, nil
}

func ClientCall(c *rpc.Client, serviceMethod string, args interface{}, reply interface{}) error{
	err := c.Call(serviceMethod, args, reply)
	if err != nil{
		for i := 1; i < MAX_RETRY_TIME; i++{
			client = nil
			c, err = GetRpcClient()
			if err != nil{
				continue
			}

			err = c.Call(serviceMethod, args, reply)
			if err == nil {
				break
			}
		}
	}
	return err
}
