//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// EvaluationItemScore 评审表项成绩模型
type EvaluationItemScore struct {
	gorm.Model
	ScoringItemID int    `gorm:"type:int;not null"`
	TeamID        int    `gorm:"type:int;not null"`
	UID           string `gorm:"type:varchar(9);not null"`
	Grade         int    `gorm:"type:int;not null"`
}

// CreateEvaluationItemScore 创建评审表项成绩
func (Repo *Repository) CreateEvaluationItemScore(evaluationItemScore EvaluationItemScore) (EvaluationItemScore, error) {
	result := Repo.DB.Create(&evaluationItemScore)
	if result.Error != nil {
		return EvaluationItemScore{}, result.Error
	}
	return evaluationItemScore, nil
}

// GetPair 用ID获取评审表项成绩
func (Repo *Repository) GetEvaluationItemScore(ID int) (EvaluationItemScore, error) {
	var evaluationItemScore EvaluationItemScore
	result := Repo.DB.First(&evaluationItemScore, ID)
	if result.Error != nil {
		return EvaluationItemScore{}, result.Error
	}
	return evaluationItemScore, nil
}

// DeleteEvaluationItemScore 根据ID删除评审表项成绩
func (Repo *Repository) DeleteEvaluationItemScore(ID int) error{
	result := Repo.DB.Delete(&EvaluationItemScore{}, ID)
	return result.Error
}

// UpdateEvaluationItemScore 更新评审表项成绩
func (Repo *Repository) UpdateEvaluationItemScore(ID int, grade int) error {
	result := Repo.DB.Model(&EvaluationItemScore{}).Where("id = ?", ID).Update("grade", grade)
	return result.Error
}

func (Repo *Repository) GetEvaluationItemScores(scoringItemID int, teamID int) ([]EvaluationItemScore, error) {
	var scores []EvaluationItemScore
	result := Repo.DB.Where("ScoringItemID = ? AND TeamID = ?", scoringItemID, teamID).Find(&scores)
	if result.Error != nil {
		return []EvaluationItemScore{}, result.Error
	}
	return scores, nil
}