package interfaces

import (
	"reflect"
	"fmt"
	"errors"
)

var storeMap = make(map[string]*RpcCall)

type RpcCall struct {
	Name      string
	Args   interface{}
	Reply interface{}

	ArgFieldName []string
}


func Register(name string, args interface{}, reply interface{}) {
	argsValue := reflect.ValueOf(args)
	replyValue := reflect.ValueOf(reply)
	if argsValue.Kind() != reflect.Ptr || replyValue.Kind() != reflect.Ptr{
		panic(errors.New(fmt.Sprintf("%s args and reply not ptr..", name)))
	}

	cal := new(RpcCall)
	cal.Name = name
	cal.Args = args
	cal.Reply = reply

	typeof := argsValue.Elem().Type()
	for i := 0; i < typeof.NumField(); i++{
		fieldI := typeof.Field(i)
		cal.ArgFieldName = append(cal.ArgFieldName, fieldI.Name)
	}
	storeMap[name] = cal
}

func GetRpcCall(name string) *RpcCall{
	return storeMap[name]
}
