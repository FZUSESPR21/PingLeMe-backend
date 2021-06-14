//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"errors"
	"gorm.io/gorm"
)

// EvaluationItemScore 评审表项成绩模型
type EvaluationItemScore struct {
	gorm.Model
	TableID uint    `gorm:"not null"`
	ItemID  uint    `gorm:"not null"`
	TeamID  uint    `gorm:"not null"`
	Score   float32 `gorm:"not null"`
}

type EvaluationItemScoreRepositoryInterface interface {
	CreateEvaluationItemScore(evaluationItemScore EvaluationItemScore) (EvaluationItemScore, error)
	GetEvaluationItemScore(ID int) (EvaluationItemScore, error)
	UpdateEvaluationItemScore(ID int, score int) error
	GetEvaluationItemScores(scoringItemID int, teamID int) ([]EvaluationItemScore, error)
	InitEvaluationItems(teamID, tableID uint) error
}

// CreateEvaluationItemScores 创建评审表项成绩
func (Repo *Repository) CreateEvaluationItemScores(items EvaluationItemScore) error {
	result := Repo.DB.Model(&EvaluationItemScore{}).Select("Score").Updates(items)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// InitEvaluationItems 初始化成绩
func (Repo *Repository) InitEvaluationItems(teamID, tableID uint) error {
	// 初始化总分
	if result := Repo.DB.Where("table_id = ? AND team_id = ?", tableID, teamID).First(&FinalEvaluationTableScore{}); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		finalScore := FinalEvaluationTableScore{
			TableID:    tableID,
			TeamID:     teamID,
			FinalScore: 0,
		}
		Repo.DB.Create(&finalScore)
	}

	var scoreItem EvaluationItemScore
	if r := Repo.DB.Where("table_id = ? AND team_id = ?", tableID, teamID).First(&scoreItem); errors.Is(r.Error, gorm.ErrRecordNotFound) {
		var table []EvaluationTableItem
		result := Repo.DB.Where("evaluation_table_id = ?", tableID).Find(&table)
		if result.Error != nil {
			return result.Error
		}
		scores := make([]EvaluationItemScore, 0)
		for _, item := range table {
			scores = append(scores, EvaluationItemScore{
				TableID: tableID,
				ItemID:  item.ID,
				TeamID:  teamID,
				Score:   0,
			})
		}
		result = Repo.DB.Create(scores)
		return result.Error
	}
	return nil

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

// UpdateEvaluationItemScore 更新评审表项成绩
func (Repo *Repository) UpdateEvaluationItemScore(ID int, score int) error {
	result := Repo.DB.Model(&EvaluationItemScore{}).Where("id = ?", ID).Update("score", score)
	return result.Error
}

// GetEvaluationItemScores 获取
func (Repo *Repository) GetEvaluationItemScores(scoringItemID int, teamID int) ([]EvaluationItemScore, error) {
	var scores []EvaluationItemScore
	result := Repo.DB.Where("item_id = ? AND team_id = ?", scoringItemID, teamID).Find(&scores)
	if result.Error != nil {
		return []EvaluationItemScore{}, result.Error
	}
	return scores, nil
}
