//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Performance 绩效模型
type Performance struct {
	gorm.Model
	HomeworkID int `gorm:"type:int;not null"`
	StudentID  int `gorm:"type:int;not null"`
	Percentage int `gorm:"type:int;not null"`
}

type PerformanceRepositoryInterface interface {
	GetPerformance(ID interface{}) (Performance, error)
	SetPerformance(performance Performance) (int64, error)
	SetPercentageByID(ID interface{}, percentage int) error
}

func (Repo *Repository) GetPerformance(ID interface{}) (Performance, error) {
	var performance Performance
	result := Repo.DB.First(&performance, ID)
	return performance, result.Error
}

func (Repo *Repository) SetPerformance(performance Performance) (int64, error) {
	result := Repo.DB.Create(&performance)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) SetPercentageByID(ID interface{}, percentage int) error {
	var performance Performance
	result := Repo.DB.Model(&performance).Where("ID = ?", ID).Update("Percentage", percentage)
	return result.Error
}
