//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type EvaluationTableService struct {
	model.EvaluationTableRepositoryInterface
	TableName  string                `json:"table_name" binding:"required"`
	HomeworkID uint                  `json:"homework_id" binding:"required"`
	TeamID     uint                  `json:"team_id" binding:"required"`
	TableItems []EvaluationTableItem `json:"table_items"`
}

type EvaluationTableDetailService struct {
	model.EvaluationTableRepositoryInterface
}

type EvaluationTableItem struct {
	Content       string                `json:"content"`
	Score         int                   `json:"score" binding:"gte=0"`
	ActualScore   int                   `json:"actual_score" binding:"gte=0"`
	ChildrenItems []EvaluationTableItem `json:"children_items"`
}

// CreateEvaluationTable 创建评审表
func (service *EvaluationTableService) CreateEvaluationTable() serializer.Response {
	table := model.EvaluationTable{
		TableName:  service.TableName,
		HomeworkID: service.HomeworkID,
		TeamID:     service.TeamID,
	}

	table.TableItem = GetChildrenItems(service.TableItems, 1)

	err := service.SetEvaluationTable(table)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// ViewEvaluationTable 查看评审表
func (service *EvaluationTableDetailService) ViewEvaluationTable(ID uint) serializer.Response {
	table, err := service.GetEvaluationTable(ID)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Data: serializer.BuildEvaluationTable(table),
	}
}

// MarkEvaluationTable 评审表评分
func (service *EvaluationTableService) MarkEvaluationTable() {
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
