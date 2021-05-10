//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// FillInPerformanceService 填写绩效的服务
type FillInPerformanceService struct {
	model.PerformanceRepositoryInterface
	HomeworkID int `form:"homework_id" json:"homework_id" binding:"required"`
	StudentID  int `form:"student_id" json:"student_id" binding:"required"`
	Percentage int `form:"percentage" json:"percentage" binding:"required"`
}

// FillInPerformance 填写绩效函数
func (service *FillInPerformanceService) FillInPerformance() serializer.Response {
	performance := model.Performance{
		HomeworkID: service.HomeworkID,
		StudentID:  service.StudentID,
		Percentage: service.Percentage,
	}
	_, err := service.SetPerformance(performance)
	if err != nil {
		return serializer.DBErr("填写绩效失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
