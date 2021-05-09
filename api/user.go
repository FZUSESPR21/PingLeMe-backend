//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"PingLeMe-Backend/util"
	"go.uber.org/zap"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

func GetTeacherList(c *gin.Context)  {
	var service service.GetTeacherListService
	res := service.GetTeacherList()
	c.JSON(200, res)
}



func AddTeachers(c *gin.Context)  {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetTeachers(c *gin.Context)  {
	var service service.GetTeacherListService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.GetTeacherList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}