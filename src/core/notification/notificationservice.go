package notification

import (
	"ec-notification-management/src/core/models/dtos"
	notificationModel "ec-notification-management/src/core/models/notification"
	"ec-notification-management/src/core/repository"
)

type Notificationservice struct {
}

var _restfulService repository.IRestfulConstraint

func init() {
	_restfulService = &repository.RestFulService{}
}

func (notification *Notificationservice) GetMsgCount(request dtos.GetMsgCountInputDto) notificationModel.OperationResult {
	res, err := _restfulService.GetMsgCount("simple", "getmsgcount")
	if err != nil {
		return notificationModel.OperationResult{
			Successed: false,
			Data:      nil,
			Message:   err.Error(),
		}
	}
	return notificationModel.OperationResult{
		Successed: true,
		Data:      res,
		Message:   "ok",
	}
}
