//  Copyright (c) 2021 PingLeMe Team. All rights reserved.
package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"github.com/gin-gonic/gin"
)


type CreateClassService struct {
	model.ClassRepositoryInterface
	model.UserRepositoryInterface
	ClassName	string	`form:"class_name" jason:"class_name binding:"required,min=5,max=30"`
	Assistants  []AssistantService	`form:"assistant_list" jason:"assistant_list" binding:"required"`
}

type AssistantService struct {
	AssistantID		int    `form:"assistant_id" jason:"assistant_id binding:"required"`
}

func (service *CreateClassService) CreateClass(c *gin.Context, teacherID int) serializer.Response{
	class, err1 := service.AddClass(service.ClassName)
	if err1 != nil {
		return serializer.DBErr("班级创建失败", err1)
	}

	teacher, err2 := service.GetUser(teacherID)
	if err2 != nil {
		return serializer.DBErr("", err2)
	}

	err1 = class.AddTeacher(teacher)
	if err1 != nil {
		return serializer.DBErr("", err1)
	}

	for i := 0; i < len(service.Assistants) - 1; i++ {
		assistant, err3 := service.GetUser(service.Assistants[i].AssistantID)
		if err3 != nil {
			return serializer.DBErr("", err3)
		}

		err1 = class.AddTeacher(assistant)
		if err1 != nil {
			return serializer.DBErr("", err1)
		}
	}

	return serializer.Response{
		Code:		0,
		Msg:		"Success",
	}
}