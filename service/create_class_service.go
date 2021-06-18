//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// CreateClassService 创建班级的服务
type CreateClassService struct {
	model.ClassRepositoryInterface
	model.UserRepositoryInterface
	ClassName  string             `form:"class_name" json:"class_name" binding:"required,min=5,max=30"`
	TeacherID  uint               `form:"teacher_id" json:"teacher_id" binding:"required,gte=0"`
	Assistants []AssistantService `form:"assistant_list" json:"assistant_list" binding:"required"`
}

type AssistantService struct {
	AssistantID uint `form:"assistant_id" json:"assistant_id" binding:"required,gte=0"`
}

// CreateClass 创建班级函数
func (service *CreateClassService) CreateClass() serializer.Response {
	class, err1 := service.AddClass(service.ClassName)
	if err1 != nil {
		return serializer.DBErr("班级创建失败", err1)
	}

	teacher, err2 := service.GetUser(service.TeacherID)
	if err2 != nil {
		return serializer.ParamErr("", err2)
	} else if teacher.Role != model.RoleTeacher {
		return serializer.ParamErr("用户不是教师", nil)
	}

	err1 = service.AddTeacher(class, teacher)
	if err1 != nil {
		return serializer.DBErr("", err1)
	}

	for _, i := range service.Assistants {
		assistant, err3 := service.GetUser(i.AssistantID)
		if err3 != nil {
			return serializer.ParamErr("", err3)
		} else if teacher.Role != model.RoleAssistant {
			return serializer.ParamErr("用户不是助教", nil)
		}

		err1 = service.AddTeacher(class, assistant)
		if err1 != nil {
			return serializer.DBErr("分配助教失败", err1)
		}
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
