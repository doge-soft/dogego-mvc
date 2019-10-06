package controllers

import (
	"dogego-mvc/web/services"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	service := services.PingService{}

	context.String(200, service.Ping())
}
