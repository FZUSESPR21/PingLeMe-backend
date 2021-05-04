//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// CreateAssistantService 添加助教的服务
type CreateAssistantService struct {
	model.UserRepositoryInterface
	UID      string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	Name     string `form:"name" json:"name" binding:"required,min=3,max=255"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// CreateAssistant 创建助教函数
func (service *CreateAssistantService) CreateAssistant() serializer.Response {
	assistant := model.User{
		UID:            service.UID,
		Nickname:       service.Name,
		Role:           model.RoleAssistant,
	}
	err := assistant.SetPassword(service.Password)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	
	err = service.SetUser(assistant)
	if err != nil {
		return serializer.DBErr("添加助教失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
