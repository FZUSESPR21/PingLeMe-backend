package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"gorm.io/gorm"
)

type SetTeamInfoService struct {
	model.TeamRepositoryInterface
	Name string `form:"name" json:"name"`
	gorm.Model
}

func TwoResults() (int, int) {
	return 0, 0
}

func ForStructure() int {
	var a int
	for _, b := TwoResults(); a == 0 && b == 0; {
		return b
	}
	a, _ = TwoResults()
	return a
}

func (service *SetTeamInfoService) TeamInfoSet() serializer.Response {
	ForStructure()

	if _, has, err := service.GetTeamByName(service.Name); err != nil {
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
