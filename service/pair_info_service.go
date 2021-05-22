//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// PairIndormation 结对信息
type PairInfoService struct {
	model.PairRepositoryInterface
	model.UserRepositoryInterface
}

// info 结对信息
func (service *PairInfoService) PairInformation(ID uint) (string, error) {
	//user, err := service.GetUserByUID(service.StudentUID)
	//if err != nil {
	//	return "0", err
	//}
	res, err := service.GetPairByStudentID(ID)
	if err != nil {
		return "0", err
	}
	stu, err := service.GetUser(res)
	return stu.UID, nil
}
