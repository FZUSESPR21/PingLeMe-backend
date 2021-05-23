//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package main

import (
	"PingLeMe-Backend/conf"
	"PingLeMe-Backend/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":3000")
}
