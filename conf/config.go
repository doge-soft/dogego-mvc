package conf

import (
	"dogego-mvc/inits"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	// 初始化
	inits.Init()
}
