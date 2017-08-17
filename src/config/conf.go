package config

import (
	"github.com/jinzhu/configor"
	"os"
	"fmt"
	"log"
)

type AppConfig struct {
	GRpc struct{
		Addr string
	}

	Server struct{
		Addr string
	}
}

var conf AppConfig

func init(){
	err := configor.Load(&conf, "conf.yml")
	if err != nil{
		log.Println(fmt.Sprintf("load conf err %v", err))
		os.Exit(1)
	}
}

func GetConfig() *AppConfig{
	return &conf
}
