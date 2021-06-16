//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// TeamListService 获取团队列表的服务
type TeamListService struct {
	model.TeamRepositoryInterface
	ClassID int `form:"class_id" json:"class_id" binding:"required"`
}

func (service *TeamListService) GetTeamList() serializer.Response {
	teams, err := service.GetTeamsByClassID(service.ClassID)
	if err != nil {
		return serializer.DBErr("不存在团队", err)
	}

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildTeamList(teams),
		Msg:  "Success",
	}
}
