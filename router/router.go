//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package router

import (
	"PingLeMe-Backend/api"
	"PingLeMe-Backend/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		v1.GET("teacher/list",api.GetTeachers)
		v1.POST("team/create",api.CreateTeam)
		v1.POST("team/member/add",api.AddTeammate)
		v1.POST("team/member/remove",api.DeleteTeammate)
		v1.POST("user/teacher/add",api.AddTeachers)



		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.LoginRequired())
		{
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
