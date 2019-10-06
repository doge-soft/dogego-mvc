package main

import (
	_ "dogego-mvc/conf"
	"dogego-mvc/web/routers"
	"dogego-mvc/web/servers"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
