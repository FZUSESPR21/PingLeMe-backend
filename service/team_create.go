package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type CreateTeamService struct {
	model.TeamRepositoryInterface
	Team model.Team `form:"team" json:"team"`
}

func (service *CreateTeamService) CreateTeam() serializer.Response {
	if has, err := service.GetTeamByName(service.Team.Name); err != nil {
		return serializer.DBErr("数据获取错误", err)
	} else if has == 1 {
		return serializer.DBErr("创建团队失败！该名称已被使用！", err)
	}
	if has, err := service.SetTeam(service.Team); err != nil || has != 1 {
		return serializer.DBErr("数据获取错误", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "创建团队成功",
	}
}
