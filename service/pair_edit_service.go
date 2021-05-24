//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// PairEditService 填写结对信息
type PairEditService struct {
	model.PairRepositoryInterface
	model.UserRepositoryInterface
	Student1UID string `json:"Student1UID"`
	Student2UID string `json:"Student2UID"`
}

// EditPairInformation 填写结对信息
func (service *PairEditService) EditPairInformation() (int, error) {
	stu1, err := service.GetUserByUID(service.Student1UID)
	if err != nil {
		return 0, err
	}
	stu2, err := service.GetUserByUID(service.Student2UID)
	if err != nil {
		return 0, err
	}

	res, err := service.UpdatePairByStu(stu1.ID, stu2.ID)
	if err != nil {
		return 0, err
	}
	return res, nil
}
