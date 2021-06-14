package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type CreateTeamService struct {
	model.TeamRepositoryInterface
	Name          string `form:"name" json:"name"`
	GroupLeaderID int    `form:"group_leader_id" json:"group_leader_id"`
	ClassID       int    `form:"class_id" json:"class_id"`
}

func (service *CreateTeamService) CreateTeam() serializer.Response {
	//TODO 1.创建者是否已有团队 2.班级是否存在 3.队名是否重复

	has, err := service.SetTeam(model.Team{
		Name:          service.Name,
		GroupLeaderID: service.GroupLeaderID,
		ClassID:       service.ClassID,
	})
	if err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	if has != 1 {
		return serializer.DBErr("has != 1 错误", err)
	}
	//TODO 返回Team_id
	return serializer.Response{
		Code: 0,
		Msg:  "创建团队成功",
	}
}
