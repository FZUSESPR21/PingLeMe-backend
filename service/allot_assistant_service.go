//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// AllotAssistantService 设置助教班级的服务
type AllotAssistantService struct {
	model.ClassRepositoryInterface
	model.UserRepositoryInterface
	UID     string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	ClassID int    `form:"class_id" json:"class_id" binding:"required"`
}

// AllotAssistant 设置助教班级函数
func (service *AllotAssistantService) AllotAssistant() serializer.Response {
	class, err := service.GetClassByID(service.ClassID)
	if err != nil {
		return serializer.ParamErr("该班级不存在", err)
	}

	assistant, err1 := service.GetUserByUID(service.UID)
	if err1 != nil {
		return serializer.ParamErr("该助教不存在", err1)
	}

	err = service.AddTeacher(class, assistant)
	if err != nil {
		return serializer.DBErr("设置助教班级失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
