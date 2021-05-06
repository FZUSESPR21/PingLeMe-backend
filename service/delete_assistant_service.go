//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// DeleteAssistantService 删除助教的服务
type DeleteAssistantService struct {
	model.UserRepositoryInterface
	AssistantID int `form:"assistant_id" json:"assistant_id" binding:"required"`
}

// DeleteAssistant 删除助教函数
func (service *DeleteAssistantService) DeleteAssistant() serializer.Response {
	err := service.DeleteUser(service.AssistantID)
	if err != nil {
		return serializer.DBErr("删除助教失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
