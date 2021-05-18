//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"go.uber.org/zap"
)

type EvaluationTableService struct {
	model.EvaluationTableRepositoryInterface
	TableName  string                `json:"table_name" binding:"required"`
	HomeworkID uint                  `json:"homework_id" binding:"required"`
	TeamID     uint                  `json:"team_id" binding:"required"`
	TableItems []EvaluationTableItem `json:"table_items"`
}

type EvaluationTableItem struct {
	Content       string                `json:"content"`
	Score         float32               `json:"score" binding:"gte=0"`
	ActualScore   float32               `json:"actual_score" binding:"gte=0"`
	ChildrenItems []EvaluationTableItem `json:"children_items"`
}

type EvaluationTableDetailService struct {
	model.EvaluationTableRepositoryInterface
}

type EvaluationTableListService struct {
	model.EvaluationTableRepositoryInterface
	model.TeamRepositoryInterface
	TeamID     uint `json:"team_id" binding:"required;gte=0"`
	HomeworkID uint `json:"homework_id" binding:"required;gte=0"`
}

type EvaluationTableScoreInitService struct {
	model.EvaluationItemScoreRepositoryInterface
	TeamID  uint `json:"team_id" binding:"required;gte=0"`
	TableID uint `json:"table_id" binding:"required;gte=0"`
}

type EvaluationTableScoreService struct {
	model.EvaluationTableRepositoryInterface
	model.FinalTableScoreRepositoryInterface
	model.EvaluationItemScoreRepositoryInterface
	TableID                   uint                       `json:"table_id" binding:"required,gte=0"`
	TeamID                    uint                       `json:"team_id" binding:"required,gte=0"`
	EvaluationTableScoreItems []EvaluationTableScoreItem `json:"score_items" binding:"required"`
}

type EvaluationTableScoreItem struct {
	ItemID uint    `json:"item_id" binding:"required;gte=0"`
	Score  float32 `json:"score" binding:"required;"`
}

// CreateEvaluationTable 创建评审表
func (service *EvaluationTableService) CreateEvaluationTable() serializer.Response {
	table := model.EvaluationTable{
		TableName:  service.TableName,
		HomeworkID: service.HomeworkID,
		TeamID:     service.TeamID,
	}

	table.TableItem = GetChildrenItems(service.TableItems, 1)
	for index, _ := range table.TableItem {
		table.TableItem[index].Index = index
	}

	err := service.SetEvaluationTable(table)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// ViewEvaluationTableScore 查看评审表
// TODO finish this.
func (service *EvaluationTableDetailService) ViewEvaluationTableScore(ID uint) serializer.Response {
	table, err := service.GetEvaluationTable(ID)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.BuildEvaluationTableResponse(table)
}

// ViewEvaluationTable 查看评审表
func (service *EvaluationTableDetailService) ViewEvaluationTable(ID uint) serializer.Response {
	table, err := service.GetEvaluationTable(ID)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.BuildEvaluationTableResponse(table)
}

// InitEvaluationTableScore 初始化成绩
func (service *EvaluationTableScoreInitService) InitEvaluationTableScore() serializer.Response {
	err := service.InitEvaluationItems(service.TeamID, service.TableID)
	if err != nil {
		return serializer.DBErr("", err)
	} else {
		return serializer.Response{
			Code: 0,
			Msg:  "Success",
		}
	}
}

// AddEvaluationTableScore 评审表评分
func (service *EvaluationTableScoreService) AddEvaluationTableScore() serializer.Response {
	table, err := service.GetEvaluationTable(service.TableID)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	tableMap := make(map[uint]model.EvaluationTableItem, 0)
	for _, t := range table.TableItem {
		tableMap[t.ID] = t
	}

	scoreItems := make(map[uint]model.EvaluationItemScore, 0)
	for _, i := range service.EvaluationTableScoreItems {
		item, ok := tableMap[i.ItemID]
		if !ok {
			return serializer.ParamErr("添加评审表错误", nil)
		}
		if i.Score > item.Score {
			return serializer.ParamErr("添加评审表错误，评审表分数错误", nil)
		}
		scoreItems[i.ItemID] = model.EvaluationItemScore{
			ItemID: i.ItemID,
			TeamID: service.TeamID,
			Score:  i.Score,
		}
	}

	tableTree := model.TransformEvaluationTableTree(table)
	totalScore := model.GetChildrenScore(tableTree.TableItems, scoreItems)

	go service.UpdateTableScore(tableTree, totalScore)

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// UpdateTableScore 更新成绩
func (service *EvaluationTableScoreService) UpdateTableScore(tableTree model.EvaluationTableTree, totalScore float32) {
	if err := service.InitEvaluationItems(service.TeamID, service.TableID); err != nil {
		util.Log().Error("成绩初始化错误", zap.Error(err))
	}

	for _, item := range tableTree.TableItems {
		err := service.UpdateScore(item, service.TeamID)
		if err != nil {
			util.Log().Error("更新成绩错误", zap.Error(err))
		}
	}
	err := service.UpdateFinalEvaluationTableScore(service.TableID, service.TeamID, totalScore)
	if err != nil {
		util.Log().Error("", zap.Error(err))
	}
}

// GetChildrenItems 递归获取子项
func GetChildrenItems(target []EvaluationTableItem, level int) []model.EvaluationTableItem {
	items := make([]model.EvaluationTableItem, 0)
	for _, item := range target {
		items = append(items, model.EvaluationTableItem{
			Content: item.Content,
			Score:   item.Score,
			Level:   level,
		})
		if item.ChildrenItems != nil {
			items = append(items, GetChildrenItems(item.ChildrenItems, level+1)...)
		}
	}
	return items
}

// GetTableList 获取评审表列表
func (service *EvaluationTableListService) GetTableList(user model.User) serializer.Response {
	var teamID uint
	switch user.Role {
	case model.RoleStudent:
		if team, err := service.GetTeamByTeamLeader(user.ID); err != nil {
			return serializer.ParamErr("", err)
		} else {
			teamID = team.ID
		}
		if list, err := service.GetEvaluationTableList(service.HomeworkID, teamID); err != nil {
			return serializer.ServerInnerErr("", err)
		} else {
			return serializer.BuildEvaluationTableListResponse(list)
		}
	case model.RoleTeacher, model.RoleAssistant:
		if list, err := service.GetEvaluationTableList(service.HomeworkID, 0); err != nil {
			return serializer.ServerInnerErr("", err)
		} else {
			return serializer.BuildEvaluationTableListResponse(list)
		}
	default:
		return serializer.Response{
			Code:  40001,
			Msg:   "User Role Err",
		}
	}
}