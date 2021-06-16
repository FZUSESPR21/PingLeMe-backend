//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// PairEditService 填写结对信息
type PairEditService struct {
	model.PairRepositoryInterface
	model.UserRepositoryInterface
	Student1UID string `json:"Student1UID" bind:"required"`
	Student2UID string `json:"Student2UID" bind:"required"`
}

// EditPairInformation 填写结对信息
func (service *PairEditService) EditPairInformation() serializer.Response {
	stu1, err := service.GetUserByUID(service.Student1UID)
	if err != nil {
		return serializer.DBErr("Student1UID有误", err)
	}
	stu2, err := service.GetUserByUID(service.Student2UID)
	if err != nil {
		return serializer.DBErr("Student2UID有误", err)
	}

	res, err := service.UpdatePairByStu(stu1.ID, stu2.ID)
	if err != nil {
		return serializer.DBErr("更新操作错误", err)
	}
	return serializer.Response{
		Code: 0,
		Data: res,
		Msg:  "Success",
	}
}
