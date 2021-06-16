//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package router

import (
	"PingLeMe-Backend/api"
	"PingLeMe-Backend/middleware"
	"net/http"
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

	r.MaxMultipartMemory = 15 << 20

	debug := r.Group("/debug")
	debug.Use(middleware.DebugAPI())
	{
		debug.StaticFS("/file", http.Dir("./blog"))

		debug.POST("ping", api.Ping)

		debug.POST("user/add", api.DebugAddUser)

		debug.POST("homework/submit/:id", api.SubmitWorks)

		debug.POST("submit/:classID", api.SubmitWorks)

		debugAuth := debug.Group("/auth")
		debugAuth.Use(middleware.LoginRequired())
		{
			debugAuth.StaticFS("/file", http.Dir("./blog"))
			debugAuth.POST("/ping", api.Ping)
		}
	}

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户登录
		v1.POST("login", api.UserLogin)

		//v1.POST("user/assistant/add", api.AddAss)

		// 班级结对状态
		v1.GET("class/pair/status/:id", api.PairStatus)

		// 班级团队状态
		v1.GET("class/team/status/:id", api.TeamStatus)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.LoginRequired())
		{
			// 用户登出
			auth.DELETE("user/logout", api.UserLogout)

			// 当前用户信息
			auth.GET("user/me", api.UserMe)

			// 用户信息
			auth.GET("user/:id", api.UserInfo)

			// 学生结对
			auth.POST("user/pair/add", api.FillInPairInformation)

			// 创建团队
			auth.POST("team/create", api.CreateTeam)

			// 修改密码
			auth.POST("user/password/change", api.ChangePassword)

			// 增加组员
			auth.POST("team/member/add", api.AddTeammate)

			// 删除组员
			auth.POST("team/member/remove", api.DeleteTeammate)

			// 填写绩效
			auth.POST("team/performance/edit", api.FillInPerformance)

			// 教师列表
			auth.GET("teacher/list", api.GetTeachers)

			// 获取班级作业列表
			auth.GET("class/homework/list/:id", api.GetHomeworkList)

			// 查看班级学生列表
			auth.GET("class/student/list/:class_id", api.ClassStuList)

			// 查看班级助教列表
			auth.GET("class/assistant/list/detail/:class_id", api.ClassAssisList)

			// 改变学生班级
			auth.POST("class/student/move", api.EditStuClass)

			// 班级列表
			auth.GET("class/list", api.ClassList)

			// 查看作业列表
			auth.POST("homework/list", api.ViewHomework)

			// 创建作业
			auth.POST("homework/create", api.CreateHomework)

			// 作业预览
			//auth.GET("homework/detail/:id")

			// 作业静态链接
			auth.StaticFS("/homework/detail/view", http.Dir("./blog"))

			// 获取评审表
			auth.GET("evaluation-table/detail/:id", api.GetEvaluationTable)

			// 填写评审表
			auth.POST("evaluation-table/fill", api.FillEvaluationTable)

			// 创建评审表
			auth.POST("evaluation-table/create", api.CreateEvaluationTable)

			// 获取班级所有团队
			auth.GET("/class/team/list", api.GetTeamList)

			permissionImportPDF := auth.Group("")
			permissionImportPDF.Use(middleware.PermissionRequired("work_submission"))
			{
				permissionImportPDF.POST("homework/submit/:classID", api.SubmitWorks)
			}

			permissionAddStudents := auth.Group("")
			permissionAddStudents.Use(middleware.PermissionRequired("add_students"))
			{
				// 批量添加学生
				permissionAddStudents.POST("user/student/add", api.AddStudents)

				// Excel导入学生
				permissionAddStudents.POST("user/student/import", api.StudentImport)
			}

			permissionAddAssistants := auth.Group("")
			permissionAddAssistants.Use(middleware.PermissionRequired("add_assistant"))
			{
				// 批量添加助教
				permissionAddAssistants.POST("user/assistant/add", api.CreateAssistant)
			}

			permissionAddTeacher := auth.Group("")
			permissionAddTeacher.Use(middleware.PermissionRequired("add_teacher"))
			{
				// 批量添加老师
				permissionAddTeacher.POST("user/teacher/add", api.AddTeachers)
			}

			permissionClassManagement := auth.Group("class")
			permissionClassManagement.Use(middleware.PermissionRequired("class_management"))
			{
				// 创建班级
				permissionClassManagement.POST("create", api.CreateClass)

				// 设置助教班级
				permissionClassManagement.POST("assistant/add", api.AddAssistant)

				// 移除助教班级
				permissionClassManagement.POST("assistant/remove", api.RemoveAssistant)

				//开始/结束结对
				permissionClassManagement.POST("pair/toggle", api.TogglePair)

				//开始/结束组队
				permissionClassManagement.POST("team/toggle", api.ToggleTeam)

				// 教师列表
				permissionClassManagement.GET("teacher/list/all", api.GetTeacherList)

				// 助教列表
				permissionClassManagement.GET("assistant/list/all", api.GetAssistantList)
			}

			//permissionHomeworkCorrect := auth.Group("class")
			//permissionHomeworkCorrect.Use(middleware.PermissionRequired("correct_homework"))
			{
				// 提交评分结果
				//permissionHomeworkCorrect.POST("homework/correct", )
			}
		}
	}
	return r
}
