//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// PairEditService 填写结对信息
type PairEditService struct {
	model.PairRepositoryInterface
	Student1ID uint `gorm:"type:int;not null;index:studentID"`
	Student2ID uint `gorm:"type:int;index:studentID"`
}

// EditPairInformation 填写结对信息
func (service *PairEditService) EditPairInformation() (int, error) {
	res, err := service.UpdatePairByStu(service.Student1ID, service.Student2ID)
	if err != nil {
		return 0, err
	}
	return res, nil
}
