//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"gorm.io/gorm"
)

// PairInfoService 结对信息
type PairInfoService struct {
	model.PairRepositoryInterface
	model.UserRepositoryInterface
}

// info 结对信息
func (service *PairInfoService) PairInformation(ID uint) serializer.Response {
	//user, err := service.GetUserByUID(service.StudentUID)
	//if err != nil {
	//	return "0", err
	//}
	res, err := service.GetPairByStudentID(ID)
	var stu model.User

	if res == 0 {
		return serializer.DBErr("暂无结对", err)
	} else {
		stu, err = service.GetUser(res)
		if err != nil {
			return serializer.DBErr("获取队友信息错误", err)
		}
	}

	pairID, err := service.GetPairByStudentID(stu.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pairID = 0
		} else {
			pairID = 0
		}
	}

	var pair model.User
	if pairID != 0 {
		pair, err = service.GetUser(pairID)
		if err != nil {
			pair = model.User{UID: "0", UserName: ""}
		}
	}

	teamID, err2 := service.GetUserTeamID(stu)
	if err2 != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			teamID = 0
		} else {
			teamID = 0
			//return serializer.DBErr("", err2)
		}
	}
	classID, err3 := service.GetStudentClassID(stu.ID)
	if err3 != nil {
		classID = 0
	}
	return serializer.BuildStudentResponse(stu, pair.UID, pair.UserName, teamID, classID)
}
