package main

import (
	"goblog/config"
	"goblog/router"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 设置路由
	r := router.SetupRouter()

	// 启动 Gin 服务器
	r.Run()
}
