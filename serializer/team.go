//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// Team 用户序列化器
type Team struct {
	Number        int    `json:"number"`
	Name          string `json:"name"`
	GroupLeaderID int    `json:"groupLeaderID"`
	ClassID       int    `json:"classID"`
}

// BuildTeam 序列化用户
func BuildTeam(team model.Team) Team {
	return Team{
		Number:        team.Number,
		Name:          team.Name,
		GroupLeaderID: team.GroupLeaderID,
		ClassID:       team.ClassID,
	}
}

// BuildTeamResponse 序列化用户响应
func BuildTeamResponse(team model.Team) Response {
	return Response{
		Data: BuildTeam(team),
	}
}
