//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Team 团队模型
type Team struct {
	gorm.Model
	Number        int    `gorm:"type:int;not null"`
	Name          string `gorm:"type:varchar(255);not null;unique"`
	GroupLeaderID int    `gorm:"type:int;not null"`
	ClassID       int    `gorm:"type:int;not null"`
	Students      []User `gorm:"many2many:student_team;"`
}

type TeamRepositoryInterface interface {
	GetTeam(ID interface{}) (Team, error)
	SetClassNameByID(ID interface{}, name string) (int64, error)
	SetTeam(team Team) (int64, error)
	SetTeammate(ID interface{}, students []User) (int64, error)
	GetTeamByName(name string) (int64, error)
}

func (Repo *Repository) GetTeam(ID interface{}) (Team, error) {
	var team Team
	result := Repo.DB.First(&team, ID)
	return team, result.Error
}

func (Repo *Repository) GetTeamByName(name string) (int64, error) {
	var team Team
	result := Repo.DB.Where("Name = ?", name).First(&team)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) SetClassNameByID(ID interface{}, name string) (int64, error) {
	var team Team
	result := Repo.DB.Model(&team).Where("ID = ?", ID).Update("Name", name)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) SetTeam(team Team) (int64, error) {
	result := Repo.DB.Create(&team)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) SetTeammate(ID interface{}, students []User) (int64, error) {
	var team Team
	result := Repo.DB.Model(&team).Where("ID = ?", ID).Update("students", students)
	return result.RowsAffected, result.Error
}
