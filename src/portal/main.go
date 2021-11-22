package main

import (
	"ec-notification-management/src/core/globalver"
	"ec-notification-management/src/portal/controllers"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	globalver.InitConfig()
	route := gin.Default()

	v1Route := route.Group("/api/v1/notification")
	{
		v1Route.GET("/msgcount/:customerNumber", controllers.GetMsgCount)
	}

	if strings.ToLower(globalver.RuntimeInfo.EnvName) == "dev" {
		route.Run(":5000")
	} else {
		route.Run(":80")
	}
}
