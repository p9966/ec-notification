package controllers

import (
	"ec-notification-management/src/core/models/dtos"
	notificationBusiness "ec-notification-management/src/core/notification"

	"github.com/gin-gonic/gin"
)

var _notificationContraint notificationBusiness.INotificationConstraint

func init() {
	_notificationContraint = &notificationBusiness.Notificationservice{}
}

//@Tags 登陆接口
// @Summary 用户登陆接口
// @Description
// @Produce  json

// @Router /api/v1/notification/msgcount [get]
func GetMsgCount(c *gin.Context) {
	var request dtos.GetMsgCountInputDto
	c.ShouldBindUri(&request)
	result := _notificationContraint.GetMsgCount(request)

	c.JSON(200, result)
}
