package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type CreateTeamService struct {
	model.TeamRepositoryInterface
	Name          string `form:"name" json:"name"`
	GroupLeaderId int    `form:"group_leader_id" json:"group_leader_id"`
	ClassId       int    `form:"class_id" json:"class_id"`
}

func makeTeam(name string, groupLeaderId int, classId int) model.Team {
	var team model.Team
	team.Name = name
	team.GroupLeaderID = groupLeaderId
	team.ClassID = classId
	return team
}

func (service *CreateTeamService) CreateTeam() serializer.Response {
	var team = makeTeam(service.Name, service.GroupLeaderId, service.ClassId)
	//TODO 1.创建者是否已有团队 2.班级是否存在 3.队名是否重复

	has, err := service.SetTeam(team)
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
