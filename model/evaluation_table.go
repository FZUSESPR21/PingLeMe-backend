//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"errors"
	"gorm.io/gorm"
)

// EvaluationTable 评审表模型
type EvaluationTable struct {
	gorm.Model
	TableName  string `gorm:"type:varchar(255);not null"`
	HomeworkID uint   `gorm:"type:int;not null"`
	TeamID     uint   `gorm:"type:int;not null"`
	TableItem  []EvaluationTableItem
}

// EvaluationTableItem 评审表项模型
type EvaluationTableItem struct {
	gorm.Model
	EvaluationTableID uint
	Content           string  `gorm:"type:varchar(255);not null"`
	Score             float32 `gorm:"type:int;not null;default:-1"`
	Level             int     `gorm:"not null;default:0"`
	Index             int     `gorm:"not null;default:0"`
}

// EvaluationTableTree 树状评审表序列化器
type EvaluationTableTree struct {
	TableID    uint
	TableName  string
	TableItems []EvaluationTableTreeItem
}

// EvaluationTableTreeItem 评审表项序列化器
type EvaluationTableTreeItem struct {
	ItemID          uint
	Content         string
	Score           float32
	ChildTableItems []EvaluationTableTreeItem
}

type EvaluationTableRepositoryInterface interface {
	GetEvaluationTable(ID uint) (EvaluationTable, error)
	SetEvaluationTable(table EvaluationTable) error
	UpdateScore(tableTreeItem EvaluationTableTreeItem, teamID uint) error
	GetEvaluationTableList(homeworkID, teamID uint) ([]EvaluationTable, error)
}

// GetEvaluationTable 获取评审表
func (Repo *Repository) GetEvaluationTable(ID uint) (EvaluationTable, error) {
	var table EvaluationTable
	result := Repo.DB.First(&table, ID)
	if result.Error != nil {
		return EvaluationTable{}, result.Error
	}

	var items []EvaluationTableItem
	result = Repo.DB.Order("`index` Desc").Where("evaluation_table_id = ?", ID).Find(&items)
	if result.Error != nil {
		return EvaluationTable{}, result.Error
	}
	table.TableItem = items

	return table, nil
}

// GetEvaluationTableList 获取班级评审表列表
func (Repo *Repository) GetEvaluationTableList(homeworkID, teamID uint) ([]EvaluationTable, error) {
	var tables []EvaluationTable
	result := Repo.DB.Where("homework_id = ? AND team_id != ?", homeworkID, teamID).Find(&tables)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return tables, nil
	}
}

// SetEvaluationTable 保存评审表
func (Repo *Repository) SetEvaluationTable(table EvaluationTable) error {
	var t EvaluationTable
	result := Repo.DB.Where("team_id = ? AND homework_id = ?", table.TeamID, table.HomeworkID).First(&t)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result := Repo.DB.Create(&table)
			return result.Error
		} else {
			return result.Error
		}
	} else {
		Repo.RemoveEvaluationTable(table.TeamID, table.HomeworkID)
		result := Repo.DB.Create(&table)
		return result.Error
		return nil
	}
}

// RemoveEvaluationTable 删除评审表
func (Repo *Repository) RemoveEvaluationTable(teamID, homeworkID uint) {
	var table EvaluationTable
	Repo.DB.Where("team_id = ? AND homework_id = ?", teamID, homeworkID).First(&table)
	Repo.DB.Unscoped().Where("evaluation_table_id", table.ID).Delete(&EvaluationTableItem{})
	Repo.DB.Unscoped().Delete(&table)
}

// TransformEvaluationTableTree 评审表转为树状
func TransformEvaluationTableTree(table EvaluationTable) EvaluationTableTree {
	items := BuildTableTreeItems(0, len(table.TableItem)-1, table.TableItem)
	return EvaluationTableTree{
		TableID:    table.ID,
		TableName:  table.TableName,
		TableItems: items,
	}
}

// UpdateScore 更新评审表得分
func (Repo *Repository) UpdateScore(tableTreeItem EvaluationTableTreeItem, teamID uint) error {
	score := EvaluationItemScore{
		ItemID: tableTreeItem.ItemID,
		TeamID: teamID,
		Score:  tableTreeItem.Score,
	}
	if tableTreeItem.ChildTableItems != nil {
		for _, item := range tableTreeItem.ChildTableItems {
			err := Repo.UpdateScore(item, teamID)
			if err != nil {
				return err
			}
		}
	}
	result := Repo.DB.Where("item_id = ? AND team_id = ?", score.ItemID, score.TeamID).Updates(&score)
	return result.Error
}

// GetChildrenScore 计算子项得分
func GetChildrenScore(tableItems []EvaluationTableTreeItem, score map[uint]EvaluationItemScore) float32 {
	var totalScore float32
	totalScore = 0
	for _, item := range tableItems {
		if item.ChildTableItems != nil {
			item.Score = GetChildrenScore(item.ChildTableItems, score)
		}
		totalScore = totalScore + item.Score
	}
	return totalScore
}

// BuildTableTreeItems 构建树状结构子项
func BuildTableTreeItems(begin, end int, tableItems []EvaluationTableItem) []EvaluationTableTreeItem {
	level := tableItems[begin].Level

	b := -1
	e := -1
	heads := make([]EvaluationTableTreeItem, 0)
	items := make([]EvaluationTableTreeItem, 0)
	i := begin
	for i <= end {
		if tableItems[i].Level > level {
			e = i
			if b == -1 {
				b = i
			}
		}

		if tableItems[i].Level == level {
			if b != -1 {
				items = append([]EvaluationTableTreeItem{{
					ItemID:          tableItems[i].ID,
					Content:         tableItems[i].Content,
					Score:           tableItems[i].Score,
					ChildTableItems: BuildTableTreeItems(b, e, tableItems),
				}}, items...)
				b = -1
				e = -1
			} else {
				items = append([]EvaluationTableTreeItem{{
					ItemID:          tableItems[i].ID,
					Content:         tableItems[i].Content,
					Score:           tableItems[i].Score,
					ChildTableItems: nil,
				}}, items...)
			}
		}

		if tableItems[i].Level < level && tableItems[i].Level == 1 {
			childItems := make([]EvaluationTableTreeItem, len(items))
			copy(childItems, items)
			heads = append([]EvaluationTableTreeItem{{
				ItemID:          tableItems[i].ID,
				Content:         tableItems[i].Content,
				Score:           tableItems[i].Score,
				ChildTableItems: childItems,
			}}, heads...)
			items = make([]EvaluationTableTreeItem, 0)
		}
		i = i + 1
	}

	if i < len(tableItems) {
		return items
	} else {
		return heads
	}
}
