//  Copyright (c) 2021 PingLeMe Team. All rights reserved.
package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// RemoveAssistantService 移除助教的服务
type RemoveAssistantService struct {
	model.UserRepositoryInterface
	model.ClassRepositoryInterface
	UID     string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	ClassID int    `form:"class_id" json:"class_id" binding:"required"`
}

// RemoveAssistant 移除助教函数
func (service *RemoveAssistantService) RemoveAssistant() serializer.Response {
	class, err1 := service.GetClassByID(service.ClassID)
	if err1 != nil {
		return serializer.ParamErr("获取班级失败", err1)
	}

	assistant, err2 := service.GetUserByUID(service.UID)
	if err2 != nil {
		return serializer.ParamErr("获取助教失败", err2)
	}

	err3 := class.DeleteTeacher(assistant)
	if err3 != nil {
		return serializer.DBErr("移除助教失败", err3)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
