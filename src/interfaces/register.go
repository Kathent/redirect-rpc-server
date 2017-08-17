package interfaces

import (
	"queueing-entity/common-module/entity"
	"web-server/grpc-client"
	"im2app/im2app"
)

func RegisterAll() {
	Register("ConnRpcObject.ConnectResultReport",
		func() interface{}{
			return &entity.ConnectArgs{}
		},

		func() interface{}{
			return &entity.ConnectResult{}
		}, nil)

	Register("ConnRpcObject.EnQueue",
		func() interface{}{
			return &entity.QueueIOArgs{}
		},

		func() interface{}{
			return &entity.QueueIOResult{}
		}, nil)

	Register("ConnRpcObject.DeQueue",
		func() interface{}{
			return &entity.QueueIOArgs{}
		},

		func() interface{}{
			return &entity.QueueIOResult{}
		}, nil)

	Register("ConnRpcObject.AddQueue",
		func() interface{}{
			return &entity.QueueADArgs{}
		},

		func() interface{}{
			return &entity.QueueADResult{}
		}, nil)

	Register("ConnRpcObject.DeleteQueue",
		func() interface{}{
			return &entity.QueueADArgs{}
		},

		func() interface{}{
			return &entity.QueueADResult{}
		}, nil)

	Register("ConnRpcObject.CallEnd",
		func() interface{}{
			return &entity.UserCallEndArgs{}
		},

		func() interface{}{
			return &entity.UserCallEndResult{}
		}, nil)

	Register("ConnRpcObject.CheckIO",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		}, nil)

	Register("ConnRpcObject.ChangeCapacity",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		}, nil)

	Register("ConnRpcObject.ChangeQueueIds",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		}, nil)

	im2AppClient := im2app.NewIm2AppClient(grpc_client.GetConnClient())
	Register("SendMessageF",
		func() interface{} {
			return &im2app.SendMessage{}
		},
		func() interface{}{
			return &im2app.SendMessageAck{}
		}, im2AppClient)

	Register("UserOnlineF",
		func() interface{} {
			return &im2app.SendMessage{}
		},
		func() interface{}{
			return &im2app.SendMessageAck{}
		}, im2AppClient)

	Register("UserOfflineF",
		func() interface{} {
			return &im2app.SendMessage{}
		},
		func() interface{}{
			return &im2app.SendMessageAck{}
		}, im2AppClient)
}
