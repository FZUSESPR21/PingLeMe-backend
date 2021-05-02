//  Copyright (c) 2021 PingLeMe Team. All rights reserved.
package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"github.com/gin-gonic/gin"
)

// DeleteAssistantService 删除助教的服务
type DeleteAssistantService struct {
	model.UserRepositoryInterface
	ID	int `form:"uid" json:"uid" binding:"required"`
}

// DeleteAssistant 删除助教函数
func (service *DeleteAssistantService) DeleteAssistant(c *gin.Context) serializer.Response {
	err := service.DeleteUser(service.ID)
	if err != nil {
		return serializer.DBErr("删除助教失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg: "Success",
	}
}