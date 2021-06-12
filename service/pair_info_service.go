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
func (service *PairInfoService) PairInformation(ID uint) (model.User, error) {
	//user, err := service.GetUserByUID(service.StudentUID)
	//if err != nil {
	//	return "0", err
	//}
	res, err := service.GetPairByStudentID(ID)
	if err != nil {
		return model.User{}, err
	}
	stu, err := service.GetUser(res)
	if err != nil {
		return stu, err
	}
	return stu, nil
}
