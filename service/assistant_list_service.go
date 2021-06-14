//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// ClassStuList 班级助教列表
type ClassAssisList struct {
	model.ClassRepositoryInterface
}

// AssistantListOfClass 班级助教列表
func (service *ClassAssisList) AssistantListOfClass(classID int) serializer.Response {
	var assisListList []model.User
	var err error
	var has int

	if has, assisListList, err = service.GetAssisByClassName(classID); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	if has == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "目前该班级没有助教！",
		}
	}

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildAssistantListResponse(assisListList),
		Msg:  "Success",
	}
}
