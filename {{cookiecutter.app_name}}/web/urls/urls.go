package urls

import "github.com/gin-gonic/gin"

type URL struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var Urls = []URL{}
