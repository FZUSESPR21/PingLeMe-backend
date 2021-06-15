//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
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

	if res == 0 {
		return serializer.DBErr("暂无结对", err)
	} else {
		_, err = service.GetUser(res)
		if err != nil {
			return serializer.DBErr("获取队友信息错误", err)
		}
	}

	return serializer.Response{
		Code: 0,
		//Data: serializer.BuildStudentResponse(stu,"","",0),
		Msg: "Success",
	}
}
