package routers

import (
	"{{cookiecutter.app_name}}/web/controllers"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
	"os"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 设置HTML渲染器
	router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: os.Getenv("TEMPLATE_DIR"),
		ContentType: "text/html",
	})

	// 中间件, 顺序不能乱

	web := router.Group("/")
	{
		// 主页
		web.GET("/", controllers.Index)
	}

	api := router.Group("/api/v1")
	{
		// 心跳检查接口
		api.POST("/ping", controllers.Ping)
	}

	return router
}
