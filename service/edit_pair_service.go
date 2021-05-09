//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// EditPairIndormation 填写结对信息
type EditPairIndormationService struct {
	model.PairRepositoryInterface
	Student1ID int `gorm:"type:int;not null;index:studentID"`
	Student2ID int `gorm:"type:int;index:studentID"`
}

// Edit 填写结对信息
func (service *EditPairIndormationService) EditPairInformation() (int, error) {
	res, err := service.UpdatePairByStu(service.Student1ID, service.Student2ID)
	if err != nil {
		return 0, err
	}
	return res, nil
}
