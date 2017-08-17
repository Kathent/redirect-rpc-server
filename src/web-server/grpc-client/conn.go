package grpc_client

import (
	"google.golang.org/grpc"
	"config"
	"log"
	"fmt"
	"os"
)

var conn *grpc.ClientConn

func init(){
	dial, err := grpc.Dial(config.GetConfig().GRpc.Addr, grpc.WithInsecure())
	if err != nil {
		log.Println(fmt.Sprintf("grpc dial err %v", err))
		os.Exit(1)
	}

	conn = dial
}

func GetConnClient() *grpc.ClientConn{
	return conn
}


