//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// ClassStuList 班级学生列表
type ClassStuList struct {
	model.ClassRepositoryInterface
}

// StuListOfClass 班级学生列表
func (service *ClassStuList) StuListOfClass(classID int) serializer.Response {
	var studentList []model.User
	var err error
	var has int

	if has, studentList, err = service.GetStusByClassName(classID); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	if has == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "目前该班级没有学生！",
		}
	}

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildStudentListResponse(studentList),
		Msg:  "Success",
	}
}
