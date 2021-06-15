//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// EditStudentClassService 改变学生班级
type EditStudentClassService struct {
	model.ClassRepositoryInterface
	StudentID int `json:"uid" bind:"required"`
	NewClass  int `json:"newClass" bind:"required"`
}

// EditStudentClass 改变学生班级
func (service *EditStudentClassService) EditStudentClass() serializer.Response {
	err := service.EditStuClass(service.StudentID, service.NewClass)
	if err != nil {
		return serializer.DBErr("数据修改错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
