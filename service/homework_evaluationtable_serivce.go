package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"errors"
	"gorm.io/gorm"
)

type HomeworkEvaluationTableService struct {
	model.EvaluationTableRepositoryInterface
	model.HomeworkRepositoryInterface
	model.TeamRepositoryInterface
}

type TableList struct {
	TeamID    uint   `json:"team_id"`
	TeamName  string `json:"team_name"`
	TableID   uint   `json:"table_id"`
	TableName string `json:"table_name"`
}

func (service *HomeworkEvaluationTableService) GetHomeworkEvaluationTableList(homeworkID uint) serializer.Response {
	tableList, err := service.GetEvaluationTableListByHomeworkID(homeworkID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.DBErr("没有评审表", err)
		} else {
			return serializer.DBErr("", err)
		}
	}

	list := make([]TableList, 0)

	for _, item := range tableList {
		team, err := service.GetTeam(item.TeamID)
		if err != nil {
			util.Log().Error(err.Error())
			continue
		}
		list = append(list, TableList{
			TeamID:    item.TeamID,
			TeamName:  team.Name,
			TableID:   item.ID,
			TableName: item.TableName,
		})
	}

	return serializer.Response{
		Code: 0,
		Data: list,
		Msg:  "Success",
	}
}
