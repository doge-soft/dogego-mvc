package routers

import (
	"dogego-mvc/web/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 中间件, 顺序不能乱

	web := router.Group("/")
	{
		// 心跳检查接口
		web.POST("/ping", controllers.Ping)
	}

	return router
}
