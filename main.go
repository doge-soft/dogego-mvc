package main

import (
	_ "{{cookiecutter.app_name}}/conf"
	"{{cookiecutter.app_name}}/web/routers"
	"{{cookiecutter.app_name}}/web/servers"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
