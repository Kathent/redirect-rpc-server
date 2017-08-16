package interfaces

import (
	"reflect"
	"fmt"
	"errors"
	"utils"
)

var storeMap = make(map[string]RpcCall)

type RpcCall interface {
	//发出rpc调用
	invokeRpcCall(args interface{}) interface{}
	//初始化操作
	init(name string, args func() interface{}, reply func() interface{}, extra interface{})
}

type GRpcCall struct {
	NormalRpcCall
	client interface{}
}

func (GRpcCall) invokeRpcCall(arg interface{}) interface{}{
	return nil
}

type NormalRpcCall struct {
	Name     string
	ArgsGen  func() interface{}
	ReplyGen func() interface{}

	ArgFieldName []string
}

func (cal *GRpcCall) init(name string, args func() interface{}, reply func() interface{}, extra interface{}){
	cal.NormalRpcCall.init(name, args, reply, extra)
	cal.client = extra
}

func (NormalRpcCall) invokeRpcCall(arg interface{}) interface{}{
	return nil
}

func (cal *NormalRpcCall) init(name string, args func() interface{}, reply func() interface{}, extra interface{}){
	cal.Name = name
	cal.ArgsGen = args
	cal.ReplyGen = reply

	typeof := reflect.ValueOf(args()).Elem().Type()
	cal.ArgFieldName = addArgFieldName(typeof)
}

func Register(name string, args func() interface{}, reply func() interface{}, client interface{}) {
	argsValue := reflect.ValueOf(args())
	replyValue := reflect.ValueOf(reply())
	if argsValue.Kind() != reflect.Ptr || replyValue.Kind() != reflect.Ptr {
		panic(errors.New(fmt.Sprintf("%s args and reply not ptr..", name)))
	}

	var call RpcCall
	if client != nil{
		gRpc := new(GRpcCall)
		call = gRpc
	}else {
		nRpc := new(NormalRpcCall)
		call = nRpc
	}

	call.init(name, args, reply, client)
}

func addArgFieldName(typeOf reflect.Type) []string {
	result := make([]string, 0, 10)
	for i := 0; i < typeOf.NumField(); i++ {
		fieldI := typeOf.Field(i)

		fieldIType := fieldI.Type
		switch fieldIType.Kind() {
		//case reflect.Array:
		//case reflect.Slice:
		//	result = append(result, fieldI.Name)
		//	break
		//case reflect.Map:
		//	tmp := addArgFieldName(fieldIType.Elem())
		//	result = utils.Merge(result, tmp)
		//	break
		case reflect.Struct:
			tmp := addArgFieldName(fieldIType)
			result = utils.Merge(result, tmp)
			break
		default:
			result = append(result, fieldI.Name)
			break
		}
	}
	return result
}

func GetRpcCall(name string) RpcCall {
	return storeMap[name]
}

func GetAllCalls() map[string]RpcCall {
	return storeMap
}
