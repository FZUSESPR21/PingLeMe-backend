//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	model.UserRepositoryInterface
	UID      string `form:"uid" json:"uid" binding:"required,min=4,max=18"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	err := s.Save()
	if err != nil {
		util.Log().Error("设置Session错误", zap.Error(err))
	}
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	user, err := service.GetUserByUID(service.UID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	util.Log().Error(strconv.FormatBool(user.CheckPassword(service.Password)))
	if !user.CheckPassword(service.Password) {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
