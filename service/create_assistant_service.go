//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// CreateAssistantService 添加助教的服务
type CreateAssistantService struct {
	model.UserRepositoryInterface
	Assistants []Assistant
}

type Assistant struct {
	UID      string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	Name     string `form:"name" json:"name" binding:"required,min=3,max=255"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// CreateAssistant 创建助教函数
func (service *CreateAssistantService) CreateAssistant() serializer.Response {
	assistants := make([]model.User, len(service.Assistants))
	i := 0
	for _, a:= range service.Assistants{
		assistant := model.User{
			UID:      a.UID,
			UserName: a.Name,
			Role:     model.RoleAssistant,
		}
		err := assistant.SetPassword(a.Password)
		if err != nil {
			return serializer.ParamErr("", err)
		}
		assistants[i] = assistant
	}

	err := service.SetUsers(assistants)
	if err != nil {
		return serializer.DBErr("添加助教失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
