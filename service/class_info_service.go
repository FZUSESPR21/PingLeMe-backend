//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"gorm.io/gorm"
)

type ClassInfoService struct {
	model.ClassRepositoryInterface
	Name string `form:"class_name" json:"class_name"`
}

//TODO 接口缺失 form:"class_name"
func (service *ClassInfoService) SetClassInfo() serializer.Response {
	class, err := service.GetClassByName(service.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.DBErr("班级名称已存在！", err)
		} else {
			return serializer.DBErr("数据获取错误", err)
		}
	}

	if err := service.UpdateClassName(class, service.Name); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
