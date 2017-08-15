package interfaces

import (
	"queueing-entity/common-module/entity"
)

func RegisterAll() {
	Register("ConnRpcObject.ConnectResultReport",
		func() interface{}{
			return &entity.ConnectArgs{}
		},

		func() interface{}{
			return &entity.ConnectResult{}
		})

	Register("ConnRpcObject.EnQueue",
		func() interface{}{
			return &entity.QueueIOArgs{}
		},

		func() interface{}{
			return &entity.QueueIOResult{}
		})

	Register("ConnRpcObject.DeQueue",
		func() interface{}{
			return &entity.QueueIOArgs{}
		},

		func() interface{}{
			return &entity.QueueIOResult{}
		})

	Register("ConnRpcObject.AddQueue",
		func() interface{}{
			return &entity.QueueADArgs{}
		},

		func() interface{}{
			return &entity.QueueADResult{}
		})

	Register("ConnRpcObject.DeleteQueue",
		func() interface{}{
			return &entity.QueueADArgs{}
		},

		func() interface{}{
			return &entity.QueueADResult{}
		})

	Register("ConnRpcObject.CallEnd",
		func() interface{}{
			return &entity.UserCallEndArgs{}
		},

		func() interface{}{
			return &entity.UserCallEndResult{}
		})

	Register("ConnRpcObject.CheckIO",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		})

	Register("ConnRpcObject.ChangeCapacity",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		})

	Register("ConnRpcObject.ChangeQueueIds",
		func() interface{}{
			return &entity.Seat{}
		},

		func() interface{}{
			return new(bool)
		})
}
