package model

import "gorm.io/gorm"

type FinalEvaluationTableScore struct {
	gorm.Model
	TableID    uint `gorm:"not null;"`
	FinalScore uint
}
