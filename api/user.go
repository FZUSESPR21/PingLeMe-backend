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
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
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
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

func GetTeacherList(c *gin.Context) {
	var service service.GetTeacherListService
	res := service.GetTeacherList()
	c.JSON(200, res)
}

func AddTeachers(c *gin.Context) {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher(false)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func AddAss(c *gin.Context) {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher(true)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetTeachers(c *gin.Context) {
	var service service.GetTeacherListService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.GetTeacherList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
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
