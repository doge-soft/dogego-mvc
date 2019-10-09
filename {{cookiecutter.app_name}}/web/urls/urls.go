package urls

import (
	"github.com/gin-gonic/gin"
	"{{cookiecutter.app_name}}/web/controllers"
)

type URL struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var Urls = []URL{
	URL{"/", "GET", controllers.Index},
}
