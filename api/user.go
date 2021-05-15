//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"PingLeMe-Backend/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const StudentImportFileDst = "./.student_import/"

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	err := s.Save()
	if err != nil {
		util.Log().Error("保存Session错误", zap.Error(err))
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

func UserMe(c *gin.Context) {
	var service service.UserInfoService
	service.PairRepositoryInterface = &model.Repo
	service.UserRepositoryInterface = &model.Repo
	res := service.Information(CurrentUser(c).ID)
	c.JSON(http.StatusOK, res)
}

// UserInfo 用户信息接口
func UserInfo(c *gin.Context) {
	var service service.UserInfoService
	service.PairRepositoryInterface = &model.Repo
	service.UserRepositoryInterface = &model.Repo
	userID := c.DefaultQuery("user", "-1")
	if userID == "-1" {
		// 404
	} else {
		user, err := strconv.Atoi(userID)
		if err != nil {
			res := serializer.ParamErr("", err)
			c.JSON(http.StatusOK, res)
		} else if user < 0 {
			res := serializer.ParamErr("用户ID错误", nil)
			c.JSON(http.StatusOK, res)
		} else {
			res := service.Information(uint(user))
			c.JSON(http.StatusOK, res)
		}
	}
}

func GetTeacherList(c *gin.Context) {
	var service service.TeacherListService
	res := service.GetTeacherList()
	c.JSON(http.StatusOK, res)
}

func AddTeachers(c *gin.Context) {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func GetTeachers(c *gin.Context) {
	var service service.TeacherListService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.GetTeacherList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// StudentImport 文件导入学生（Excel）
func StudentImport(c *gin.Context) {
	file, _ := c.FormFile("file")
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(int(CurrentUser(c).ID)))
	builder.WriteString(time.Now().String())
	builder.WriteString(util.RandStringRunes(5))
	file.Filename = builder.String()

	err := c.SaveUploadedFile(file, StudentImportFileDst)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  serializer.CodeInnerError,
			Error: err.Error(),
		})
	}

	var service service.StudentImportService
	res := service.Import(StudentImportFileDst + file.Filename)
	c.JSON(http.StatusOK, res)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var service service.ChangePasswordService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.ChangePassword()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// AddStudents 批量添加学生
func AddStudents(c *gin.Context) {
	var service service.AddStudentsService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		service.ClassRepositoryInterface = &model.Repo
		res := service.AddStudents()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
