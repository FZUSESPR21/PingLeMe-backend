package model

import "gorm.io/gorm"

// AssistantScored 助教评分记录模型
type AssistantScored struct {
	gorm.Model
	HomeworkID  	uint   `gorm:"type:int;not null"`
	ScorekeeperID 	uint   `gorm:"type:int;not null"`
	AssistantID 	uint   `gorm:"type:int;not null"`
}

// AddAssistantScored 增加助教评分记录
func (Repo *Repository) AddAssistantScored(scored AssistantScored) error {
	result := Repo.DB.Create(&scored)
	return result.Error
}
