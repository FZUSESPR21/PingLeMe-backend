package serializer

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Number        int    `gorm:"type:int;not null"`
	Name          string `gorm:"type:varchar(255);not null;unique"`
	GroupLeaderID int    `gorm:"type:int;not null"`
	ClassID       int    `gorm:"type:int;not null"`
	Students      []User `gorm:"many2many:student_team;"`
}


