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
		//v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.LoginRequired())
		{
			// 用户登出
			auth.DELETE("user/logout", api.UserLogout)

			// 当前用户信息
			v1.GET("user/me", api.UserMe)

			// 用户信息
			v1.GET("user/:id", api.UserInfo)

			// 结对队友信息
			v1.GET("user/pair/:id", api.PairInfo)

			// 学生结对
			v1.POST("user/pair/edit", api.FillInPairInformation)

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

			// 教师列表
			v1.GET("teacher/list", api.GetTeachers)

			// 获取班级作业列表
			v1.GET("class/homework/list/:id", api.GetHomeworkList)

			// 查看班级学生列表
			v1.GET("class/student/list/:class_id", api.ClassStuList)

			// 查看班级助教列表
			v1.GET("class/assistant/list/:class_id", api.ClassAssisList)

			// 获取班级作业列表
			v1.GET("class/homework/list/", api.GetHomeworkList)

			// 改变学生班级
			v1.POST("class/student/move", api.EditStuClass)

			// 班级列表
			//v1.GET("class/list", )

			// 查看作业列表
			v1.POST("homework/list", api.ViewHomework)

			// 创建作业
			v1.POST("homework/create", api.CreateHomework)

			// 作业预览
			//v1.GET("homework/detail/:id")

			// 获取评审表
			v1.GET("evaluation-table/detail/:id", api.GetEvaluationTable)

			// 填写评审表
			v1.POST("evaluation-table/fill", api.FillEvaluationTable)

			// 创建评审表
			v1.POST("evaluation-table/create", api.CreateEvaluationTable)

			permissionAddStudents := v1.Group("")
			permissionAddStudents.Use(middleware.PermissionRequired("add_students"))
			{
				// 批量添加学生
				v1.POST("user/student/add", api.AddStudents)
			}


			permissionAddAssistants := v1.Group("")
			permissionAddAssistants.Use(middleware.PermissionRequired("add_assistant"))
			{
				// 批量添加助教
				v1.POST("user/assistant/add", api.CreateAssistant)
			}

			// 查看作业列表
			v1.POST("homework/list", api.ViewHomeworkList)

			permissionAddTeacher := v1.Group("")
			permissionAddTeacher.Use(middleware.PermissionRequired("add_teacher"))
			{
				// 批量添加老师
				v1.POST("user/teacher/add", api.AddTeachers)
			}

			permissionClassManagement := v1.Group("class")
			permissionClassManagement.Use(middleware.PermissionRequired("class_management"))
			{
				// 创建班级
				v1.POST("create", api.CreateClass)

				// 设置助教班级
				v1.POST("assistant/add", api.AddAssistant)


				// 移除助教班级
				v1.POST("assistant/remove", api.RemoveAssistant)

				//开始/结束结对
				v1.POST("pair/toggle", api.TogglePair)

				//开始/结束组队
				v1.POST("team/toggle", api.ToggleTeam)
			}

			// 获取评审表
			v1.GET("evaluation-table/detail/:id", api.GetEvaluationTable)

			// 填写评审表
			v1.POST("evaluation-table/fill", api.FillEvaluationTable)

			// 创建评审表
			v1.POST("evaluation-table/create", api.CreateEvaluationTable)


			permissionHomeworkCorrect := v1.Group("class")
			permissionHomeworkCorrect.Use(middleware.PermissionRequired("correct_homework"))
			{
				// 提交评分结果
				//v1.POST("homework//correct", )
			}
		}
	}
	return r
}
