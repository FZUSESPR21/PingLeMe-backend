//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
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
	Content           string `gorm:"type:varchar(255);not null"`
	Score             int    `gorm:"type:int;not null;default:-1"`
	Level             int    `gorm:"not null;default:0"`
}

type EvaluationTableRepositoryInterface interface {
	GetEvaluationTable(ID uint) (EvaluationTable, error)
	SetEvaluationTable(table EvaluationTable) error
}

// GetEvaluationTable 获取评审表
func (Repo *Repository) GetEvaluationTable(ID uint) (EvaluationTable, error) {
	var table EvaluationTable
	result := Repo.DB.First(&table, ID)
	if result.Error != nil {
		return EvaluationTable{}, result.Error
	}

	var items []EvaluationTableItem
	result = Repo.DB.Order("index desc").Where("evaluation_table_id = ?", ID).Find(&items)
	if result.Error != nil {
		return EvaluationTable{}, result.Error
	}
	table.TableItem = items

	return table, nil
}

// SetEvaluationTable 保存评审表
func (Repo *Repository) SetEvaluationTable(table EvaluationTable) error {
	result := Repo.DB.Create(&table)
	return result.Error
}
