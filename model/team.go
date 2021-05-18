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

//tpye Student_team struct {
//	team
//}

type TeamRepositoryInterface interface {
	GetTeam(ID interface{}) (Team, error)
	SetClassNameByID(ID interface{}, name string) (int64, error)
	SetTeam(team Team) (int64, error)
	SetTeammate(number int, students []User) (int64, error)
	GetTeamByName(name string) (int64, error)
	GetTeamByNumber(number int) (Team, int64, error)
	TestFunc() int
	AddTeammateByID(uid int, teamID int) (int64, error)
	DeleteTeammateByID(uid int) (int64, error)
}

func (Repo *Repository) TestFunc() int {
	return 1
}

func (Repo *Repository) GetTeam(ID interface{}) (Team, error) {
	var team Team
	result := Repo.DB.First(&team, ID)
	return team, result.Error
}

func (Repo *Repository) GetTeamByNumber(number int) (Team, int64, error) {
	var team Team
	result := Repo.DB.Where("number = ?", number).First(&team)
	return team, result.RowsAffected, result.Error
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

func (Repo *Repository) SetTeammate(number int, students []User) (int64, error) {
	var team Team
	result := Repo.DB.Model(&team).Where("Number = ?", number).Update("students", students)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) AddTeammateByID(uid int, teamID int) (int64, error) {
	//TODO 一大堆判断
	result := Repo.DB.Exec("INSERT INTO student_team VALUES (?, ?)", teamID, uid)
	return result.RowsAffected, result.Error
}

func (Repo *Repository) DeleteTeammateByID(uid int) (int64, error) {
	//TODO 一大堆判断
	result := Repo.DB.Exec("delete from student_team where user_id = ?", uid)
	return result.RowsAffected, result.Error
}
