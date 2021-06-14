//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/cache"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type ToggleGroupService struct {
}

func (service *ToggleGroupService) ToggleGroup(classID uint, duration time.Duration, str string) serializer.Response {
	val, err := cache.Get(strconv.Itoa(int(classID)), str)
	if errors.Is(err, redis.Nil) || val != "true" {
		ok, err := cache.SetNX(strconv.Itoa(int(classID))+str, "true", duration)
		if err != nil {
			util.Log().Error(zap.Error(err).String)
			return serializer.ServerInnerErr("Redis error:", err)
		} else if !ok {
			util.Log().Error("Redis SetNX get false.")
			return serializer.ServerInnerErr("Redis SetNX get false.", err)
		}
	} else if err == nil {
		_, err := cache.Del(strconv.Itoa(int(classID)) + str)
		if err != nil {
			util.Log().Error(zap.Error(err).String)
			return serializer.ServerInnerErr("Redis error:", err)
		}
	} else {
		util.Log().Error(zap.Error(err).String)
		return serializer.ServerInnerErr("Redis error:", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
