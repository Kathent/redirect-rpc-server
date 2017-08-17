package interfaces

import (
	"reflect"
	"fmt"
	"errors"
	"utils"
	"web-server/rpc-client"
	"google.golang.org/grpc"
	"context"
	"web-server/grpc-client"
	"strings"
)

var storeMap = make(map[string]RpcCall)

type RpcCall interface {
	//发出rpc调用
	InvokeRpcCall(args interface{}, reply interface{}) error
	//初始化操作
	init(name string, args func() interface{}, reply func() interface{}, extra interface{}) error
	//获取Rpc基本数据
	GetBaseRpcCall() *BaseRpcCall
}

type BaseRpcCall struct {
	Name     string
	ArgsGen  func() interface{}
	ReplyGen func() interface{}
	ArgFieldName []string
}

func NewBaseRpc(name string, args func() interface{}, reply func() interface{}) *BaseRpcCall{
	res := &BaseRpcCall{
		Name: name,
		ArgsGen: args,
		ReplyGen: reply,
	}

	typeof := reflect.ValueOf(args()).Elem().Type()
	res.ArgFieldName = addArgFieldName(typeof)
	return res
}

type GRpcCall struct {
	*BaseRpcCall
	GRpcName string
}

func (cal *GRpcCall) InvokeRpcCall(arg interface{}, reply interface{}) error{
	return grpc.Invoke(context.Background(), cal.GRpcName, arg, reply, grpc_client.GetConnClient())
}

func (cal *GRpcCall) init(name string, args func() interface{}, reply func() interface{}, extra interface{}) error{
	structValue := reflect.ValueOf(extra).Elem()
	pkgName := structValue.Type().PkgPath()
	pkgName = strings.Replace(pkgName, "/", ".", -1)
	structName := structValue.Type().Name()

	interfaceMapName := fmt.Sprintf("%s.%s", structName, name)
	gRpcName := fmt.Sprintf("/%s/%s", pkgName, name)

	cal.BaseRpcCall = NewBaseRpc(interfaceMapName, args, reply)
	cal.GRpcName = gRpcName
	return nil
}

func (cal *GRpcCall) GetBaseRpcCall() *BaseRpcCall {
	return cal.BaseRpcCall
}

type NormalRpcCall struct {
	*BaseRpcCall
}

func (cal *NormalRpcCall) InvokeRpcCall(arg interface{}, reply interface{}) error{
	client, err := rpc_client.GetRpcClient()
	if err != nil{
		return err
	}
	err = rpc_client.ClientCall(client, cal.BaseRpcCall.Name, arg, reply)
	if err != nil{
		return err
	}
	return nil
}

func (cal *NormalRpcCall) init(name string, args func() interface{}, reply func() interface{}, extra interface{}) error{
	cal.BaseRpcCall = NewBaseRpc(name, args, reply)
	return nil
}

func (cal *NormalRpcCall) GetBaseRpcCall() *BaseRpcCall {
	return cal.BaseRpcCall
}

func Register(name string, args func() interface{}, reply func() interface{}, client interface{}){
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

	err := call.init(name, args, reply, client)
	if err != nil{
		panic(err)
	}

	storeMap[call.GetBaseRpcCall().Name] = call
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
