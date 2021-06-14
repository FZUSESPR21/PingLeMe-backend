//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type DebugUserService struct {
	model.UserRepositoryInterface
	model.RBACRepositoryInterface
	UID      string `form:"uid" json:"uid"`
	UserName string `form:"name" json:"name"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Role     uint8  `form:"role" json:"role"`
}

func (service *DebugUserService) DebugAddUser() serializer.Response {
	user := model.User{
		UID:      service.UID,
		UserName: service.UserName,
		Role:     service.Role,
	}
	if service.Role != model.RoleTeacher &&
		service.Role != model.RoleAdmin &&
		service.Role != model.RoleAssistant &&
		service.Role != model.RoleStudent {
		return serializer.ParamErr("user role type error.", nil)
	}
	err1 := user.SetPassword(service.Password)
	if err1 != nil {
		return serializer.ParamErr("password parsing error.", err1)
	}

	err2 := service.SetUser(user)
	if err2 != nil {
		return serializer.ServerInnerErr("", err2)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
