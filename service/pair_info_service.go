//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"fmt"
)

// PairIndormation 填写结对信息
type PairInfoService struct {
	model.PairRepositoryInterface
	model.UserRepositoryInterface
	StudentID int `json:"uid"`
}

// info 结对信息
func (service *PairInfoService) PairInformation() (int, error) {
	user, err := service.GetUserByUID(fmt.Sprint(service.StudentID))

	res, err := service.GetPairByStudentID(int(service.GetUserID(user)))
	if err != nil {
		return 0, err
	}
	return res, nil
}
