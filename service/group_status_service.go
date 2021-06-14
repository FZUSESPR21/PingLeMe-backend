//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/cache"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
)

type GroupStatusService struct {
	model.ClassRepositoryInterface
}

func (service *GroupStatusService) Status(classID string, key string) serializer.Response {
	val, err := cache.Get(classID, key)
	id, err2 := strconv.Atoi(classID)
	if err2 != nil {
		return serializer.ParamErr("", err2)
	}
	_, err1 := service.GetClassByID(id)
	if err1 != nil {
		if err1 == gorm.ErrRecordNotFound {
			return serializer.ParamErr("class not exists.", nil)
		}
		return serializer.ServerInnerErr("", err)
	}
	if errors.Is(err, redis.Nil) || val != "true" {
		return serializer.BuildStatusResponse(false)
	} else {
		return serializer.BuildStatusResponse(true)
	}
}
