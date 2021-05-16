//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// EvaluationItemScore 评审表项成绩模型
type EvaluationItemScore struct {
	gorm.Model
	ItemID uint    `gorm:"type:int;not null"`
	TeamID uint    `gorm:"type:int;not null"`
	Score  float32 `gorm:"type:int;not null"`
}

type EvaluationItemScoreRepositoryInterface interface {
	CreateEvaluationItemScore(evaluationItemScore EvaluationItemScore) (EvaluationItemScore, error)
	GetEvaluationItemScore(ID int) (EvaluationItemScore, error)
	DeleteEvaluationItemScore(ID int) error
	UpdateEvaluationItemScore(ID int, score int) error
	GetEvaluationItemScores(scoringItemID int, teamID int) ([]EvaluationItemScore, error)
}

// CreateEvaluationItemScores 创建评审表项成绩
func (Repo *Repository) CreateEvaluationItemScores(items EvaluationItemScore) error {
	result := Repo.DB.Model(&EvaluationItemScore{}).Select("Score").Updates(items)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// InitEvaluationItems 初始化成绩
func (Repo *Repository) InitEvaluationItems(teamID, tableID uint) error {
	var table EvaluationTable
	result := Repo.DB.Preload("EvaluationTableItem").First(&table, tableID)
	if result.Error != nil {
		return result.Error
	}
	scores := make([]EvaluationItemScore, 0)
	for _, item := range table.TableItem {
		scores = append(scores, EvaluationItemScore{
			ItemID:      item.ID,
			TeamID: teamID,
			Score:  0,
		})
	}
	result = Repo.DB.Create(scores)
	return result.Error
}

// CreateEvaluationItemScore 创建评审表项成绩
func (Repo *Repository) CreateEvaluationItemScore(evaluationItemScore EvaluationItemScore) (EvaluationItemScore, error) {
	result := Repo.DB.Create(&evaluationItemScore)
	if result.Error != nil {
		return EvaluationItemScore{}, result.Error
	}
	return evaluationItemScore, nil
}

// GetEvaluationItemScore 用ID获取评审表项成绩
func (Repo *Repository) GetEvaluationItemScore(ID int) (EvaluationItemScore, error) {
	var evaluationItemScore EvaluationItemScore
	result := Repo.DB.First(&evaluationItemScore, ID)
	if result.Error != nil {
		return EvaluationItemScore{}, result.Error
	}
	return evaluationItemScore, nil
}

// DeleteEvaluationItemScore 根据ID删除评审表项成绩
func (Repo *Repository) DeleteEvaluationItemScore(ID int) error {
	result := Repo.DB.Delete(&EvaluationItemScore{}, ID)
	return result.Error
}

// UpdateEvaluationItemScore 更新评审表项成绩
func (Repo *Repository) UpdateEvaluationItemScore(ID int, score int) error {
	result := Repo.DB.Model(&EvaluationItemScore{}).Where("id = ?", ID).Update("score", score)
	return result.Error
}

func (Repo *Repository) GetEvaluationItemScores(scoringItemID int, teamID int) ([]EvaluationItemScore, error) {
	var scores []EvaluationItemScore
	result := Repo.DB.Where("scoring_item_id = ? AND team_id = ?", scoringItemID, teamID).Find(&scores)
	if result.Error != nil {
		return []EvaluationItemScore{}, result.Error
	}
	return scores, nil
}
