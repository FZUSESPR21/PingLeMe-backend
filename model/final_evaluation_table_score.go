package model

import "gorm.io/gorm"

type FinalEvaluationTableScore struct {
	gorm.Model
	TableID    uint `gorm:"not null;"`
	TeamID     uint `gorm:"not null;"`
	FinalScore float32
}

type FinalTableScoreRepositoryInterface interface {
	UpdateFinalEvaluationTableScore(tableID, teamID uint, score float32) error
}

// InitFinalEvaluationTableScore 初始化总成绩
func (Repo *Repository) InitFinalEvaluationTableScore(tableID, teamID uint) error {
	finalScore := FinalEvaluationTableScore{
		TableID:    tableID,
		TeamID:     teamID,
		FinalScore: 0,
	}
	result := Repo.DB.Create(&finalScore)
	return result.Error
}

// UpdateFinalEvaluationTableScore 更新成绩
func (Repo *Repository) UpdateFinalEvaluationTableScore(tableID, teamID uint, score float32) error {
	finalScore := FinalEvaluationTableScore{
		TableID:    tableID,
		TeamID:     teamID,
		FinalScore: score,
	}
	result := Repo.DB.Where("table_id = ? AND team_id = ?", tableID, teamID).Updates(&finalScore)
	return result.Error
}
