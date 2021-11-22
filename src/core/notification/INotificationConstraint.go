package notification

import (
	"ec-notification-management/src/core/models/dtos"
	"ec-notification-management/src/core/models/notification"
)

type INotificationConstraint interface {
	GetMsgCount(request dtos.GetMsgCountInputDto) notification.OperationResult
}
