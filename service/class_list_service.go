//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/cache"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"github.com/go-redis/redis/v8"
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
	for _, i := range info {
		i.PairStatus, _ = checkStatus(i.ClassID, "pair")
		i.TeamStatus, _ = checkStatus(i.ClassID, "team")
	}
	return serializer.Response{
		Code: 0,
		Data: info,
		Msg:  "Success",
	}
}

func checkStatus(classID uint, key string) (bool, error) {
	val, err := cache.Get(strconv.Itoa(int(classID)), key)
	if errors.Is(err, redis.Nil) || val != "true" {
		return false, nil
	} else if err != nil {
		return true, nil
	}
	if val == "true" {
		return true, nil
	}
	return false, nil
}
