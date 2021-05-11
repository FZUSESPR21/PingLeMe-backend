//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// HomeworkListService 获取作业列表的服务
type HomeworkListService struct {
	model.HomeworkRepositoryInterface
	model.ClassRepositoryInterface
	ClassID uint `json:"class_id" binding:"required"`
	Page    int  `json:"page"`
}

// GetHomeworkList 获取作业列表函数
func (service *HomeworkListService) GetHomeworkList() serializer.Response {
	class, err := service.GetClassByID(service.ClassID)
	if err != nil {
		return serializer.ParamErr("该班级不存在", err)
	}

	homeworks, err1 := class.GetAllHomework()
	if err1 != nil {
		return serializer.ParamErr("获取作业列表失败", err)
	}

	pages := len(homeworks) / 5
	homeworks = homeworks[(service.Page-1)*5 : (service.Page-1)*5+5]

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildHomeworkList(homeworks, pages, service.Page),
	}
}
