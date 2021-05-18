package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"gorm.io/gorm"
)

type TeamInfoService struct {
	model.TeamRepositoryInterface
	Name string `form:"name" json:"name"`
	gorm.Model
}

func (service *TeamInfoService) TeamInfoSet() serializer.Response {
	if has, err := service.GetTeamByName(service.Name); err != nil {
		return serializer.DBErr("数据获取错误", err)
	} else if has == 1 {
		return serializer.DBErr("修改失败！该名称已被使用！", err)
	}
	if has, err := service.SetClassNameByID(service.ID, service.Name); err != nil || has != 1 {
		return serializer.DBErr("数据获取错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "修改成功",
	}
}
