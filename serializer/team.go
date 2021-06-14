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

// TeamList 团队列表序列化器
type TeamList struct {
	List 	[]Team 	`json:"list"`
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

// BuildTeamList 序列化团队列表
func BuildTeamList(teamList []model.Team) TeamList {
	items := make([]Team, len(teamList))
	for i := range items {
		items[i].Number = teamList[i].Number
		items[i].Name = teamList[i].Name
		items[i].GroupLeaderID = teamList[i].GroupLeaderID
		items[i].ClassID = teamList[i].ClassID
	}

	return TeamList{
		List: items,
	}
}

// BuildTeamResponse 序列化用户响应
func BuildTeamResponse(team model.Team) Response {
	return Response{
		Data: BuildTeam(team),
	}
}
