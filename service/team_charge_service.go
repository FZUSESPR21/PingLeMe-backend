//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// TeamChargeService 组长管理团队的服务
type TeamManagementService struct {
	model.TeamRepositoryInterface
	Number        int    `form:"Number" json:"Number" binding:"required"`
	Name          string `form:"Name" json:"Name" binding:"required,min=1,max=30"`
	GroupLeaderID int    `form:"GroupLeaderID" json:"GroupLeaderID"`
	ClassID       int    `form:"ClassID" json:"ClassID" binding:"required"`
	UID           string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
}

// Create 组长创建团队
func (service *TeamManagementService) Create(c *gin.Context, team model.Team) serializer.Response {
	team.Number = service.Number
	team.Name = service.Name
	team.GroupLeaderID = service.GroupLeaderID
	team.ClassID = service.ClassID

	res, err := service.SetTeam(team)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.ParamErr("创建团队错误", nil)
		}
	} else {
		return serializer.ParamErr("创建团队错误，值无效", nil)
	}

	if res == 0 {
		return serializer.ParamErr("创建团队错误", nil)
	}

	team1, err := service.GetLastTeam()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("查找团队错误", nil)
	}

	return serializer.BuildTeamResponse(team1)
}

//Add 组长添加组员
func (service *TeamManagementService) Add(c *gin.Context, ID interface{}, user model.User) serializer.Response {
	team, err := service.GetTeam(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("创建团队错误", nil)
	}

	team.Students = append(team.Students, user)

	return serializer.BuildTeamResponse(team)
}

//Delete 组长删除组员
func (service *TeamManagementService) Delete(c *gin.Context, ID interface{}, user model.User, index int) serializer.Response {
	team, err := service.GetTeam(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("创建团队错误", nil)
	}

	team.Students = append(team.Students[:index], team.Students[index+1:]...)

	return serializer.BuildTeamResponse(team)
}

//Edit 组长修改团队信息
func (service *TeamManagementService) Edit(c *gin.Context, ID interface{}, name string) serializer.Response {
	team, err := service.GetTeam(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("获取团队信息错误", nil)
	}

	res, err1 := service.SetClassNameByID(ID, name)
	if res == 0 {
		return serializer.ParamErr("修改团队名称错误", nil)
	}

	if errors.Is(err1, gorm.ErrRecordNotFound) {
		return serializer.ParamErr("修改团队名称错误", nil)
	}

	return serializer.BuildTeamResponse(team)
}
