package interfaces

import (
	"reflect"
	"fmt"
	"errors"
	"utils"
)

var storeMap = make(map[string]*RpcCall)

type RpcCall struct {
	Name     string
	ArgsGen  func() interface{}
	ReplyGen func() interface{}

	ArgFieldName []string
}


func Register(name string, args func() interface{}, reply func() interface{}) {
	argsValue := reflect.ValueOf(args())
	replyValue := reflect.ValueOf(reply())
	if argsValue.Kind() != reflect.Ptr || replyValue.Kind() != reflect.Ptr{
		panic(errors.New(fmt.Sprintf("%s args and reply not ptr..", name)))
	}

	cal := new(RpcCall)
	cal.Name = name
	cal.ArgsGen = args
	cal.ReplyGen = reply

	typeof := argsValue.Elem().Type()
	cal.ArgFieldName = addArgFieldName(typeof)
	//fmt.Println(name, cal.ArgFieldName, len(cal.ArgFieldName))
	storeMap[name] = cal
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

func GetRpcCall(name string) *RpcCall{
	return storeMap[name]
}

func GetAllCalls() map[string]*RpcCall {
	return storeMap
}
