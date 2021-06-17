//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/cache"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"strconv"
)

type ClassListService struct {
	model.ClassRepositoryInterface
}

func (service *ClassListService) GetClassList() serializer.Response {
	info, err := service.GetClassInfoList()
	if err != nil {
		return serializer.ServerInnerErr("", err)
	}
	for index, i := range info {
		info[index].PairStatus, _ = CheckStatus(i.ClassID, "pair")
		info[index].TeamStatus, _ = CheckStatus(i.ClassID, "team")
	}
	return serializer.Response{
		Code: 0,
		Data: info,
		Msg:  "Success",
	}
}

func CheckStatus(classID uint, key string) (bool, error) {
	val, err := cache.Get(strconv.Itoa(int(classID)), key)
	if err != nil {
		util.Log().Debug("1")
		return false, err
	}
	if val == "true" {
		util.Log().Debug("2")
		return true, nil
	} else {
		util.Log().Debug("3")
		return false, nil
	}
}
