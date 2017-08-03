package interfaces

import "common-module/entity"

func RegisterAll() {
	Register("ConnRpcObject.ConnectResultReport", &entity.ConnectArgs{}, &entity.ConnectResult{})

	Register("ConnRpcObject.EnQueue", &entity.QueueIOArgs{}, &entity.QueueIOResult{})
	Register("ConnRpcObject.DeQueue", &entity.QueueIOArgs{}, &entity.QueueIOResult{})

	Register("ConnRpcObject.AddQueue", &entity.QueueADArgs{}, &entity.QueueADResult{})
	Register("ConnRpcObject.DeleteQueue", &entity.QueueADArgs{}, &entity.QueueADResult{})

	Register("ConnRpcObject.CallEnd", &entity.UserCallEndArgs{}, &entity.UserCallEndResult{})

	Register("SeatTrigger.CheckIO", &entity.Seat{}, new(bool))
	Register("SeatTrigger.ChangeCapacity", &entity.Seat{}, new(bool))
	Register("SeatTrigger.ChangeQueueIds", &entity.Seat{}, new(bool))
}
