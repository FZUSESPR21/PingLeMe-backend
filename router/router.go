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

	r.MaxMultipartMemory = 15

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 用户信息
		v1.GET("user/:id", api.UserInfo)

		// 当前用户信息
		v1.GET("user/me", api.UserMe)

		// 结对队友信息
		v1.GET("user/pair/:id", api.PairInfo)

		// 创建团队
		v1.POST("team/create", api.CreateTeam)

		// 修改密码
		v1.POST("user/password/change", api.ChangePassword)

		// 增加组员
		v1.POST("team/member/add", api.AddTeammate)

		// 删除组员
		v1.POST("team/member/remove", api.DeleteTeammate)

		// 填写绩效
		//v1.POST("team/performance/edit", )

		// 批量添加学生
		v1.POST("user/student/add", api.AddStudents)

		// 批量添加助教
		v1.POST("user/assistant/add", api.CreateAssistant)

		// 学生结对
		v1.POST("user/pair/add", api.FillInPairInformation)

		// 教师列表
		v1.GET("teacher/list", api.GetTeachers)

		// 批量添加老师
		v1.POST("user/teacher/add", api.AddTeachers)

		// 设置助教班级
		v1.POST("class/assistant/add", api.AddAssistant)

		// 获取班级作业列表
		v1.GET("class/homework/list/", api.GetHomeworkList)

		// 创建班级
		//v1.GET("class/create", )

		// 查看班级学生列表
		v1.GET("class/student/list/:class_id", api.ClassStuList)

		// 改变学生班级
		v1.POST("class/student/move", api.EditStuClass)

		// 移除助教
		v1.POST("class/assistant/remove", api.RemoveAssistant)

		// 班级列表
		//v1.GET("class/list", )

		//开始结对
		//v1.POST("class/pair/start/:class_id", )

		//结束结对
		//v1.POST("class/pair/end/:class_id", )

		//开始结对
		//v1.POST("class/team/start/:class_id", )

		//结束结对
		//v1.POST("class/team/end/:class_id", )

		// 查看作业列表
		v1.POST("homework/list", api.ViewHomeworkList)

		// 创建作业
		v1.POST("homework/create", api.CreateHomework)

		// 作业预览
		//v1.GET("homework/detail/:id")

		// 提交评分结果
		//v1.POST("homework/correct", )

		// 获取评审表
		v1.GET("evaluation-table/detail/:id", api.GetEvaluationTable)

		// 填写评审表
		v1.POST("evaluation-table/fill", api.FillEvaluationTable)

		// 创建评审表
		v1.POST("evaluation-table/create", api.CreateEvaluationTable)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.LoginRequired())
		{
			// 用户登出
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
