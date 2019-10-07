package routers

import (
	"{{cookiecutter.app_name}}/web/controllers"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
	"os"
	"{{cookiecutter.app_name}}/web/urls"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 设置HTML渲染器
	router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: os.Getenv("TEMPLATE_DIR"),
		ContentType: "text/html; charset=utf-8",
	})

	web := router.Group("/")
	{
		for _, url := range urls.Urls {
			web.Handle(url.Method, url.Path, url.Handler)
		}
	}

	api := router.Group("/api/v1")
	{
		// 心跳检查接口
		api.POST("/ping", controllers.Ping)
	}

	return router
}
