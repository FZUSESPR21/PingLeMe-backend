//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// PersonalBlogScore 个人博客成绩模型
type PersonalBlogScore struct {
	gorm.Model
	ScoringItemID int `gorm:"type:int;not null"`
	ScorekeeperID int `gorm:"type:int;not null"`
	Grade         int `gorm:"type:int;not null"`
}

// TeamBlogScore 团队博客成绩模型
type TeamBlogScore struct {
	gorm.Model
	ScoringItemID int `gorm:"type:int;not null"`
	ScorekeeperID int `gorm:"type:int;not null"`
	Grade         int `gorm:"type:int;not null"`
}

type BlogScoreRepositoryInterface interface {
	SetPersonalBlogScore(personalBlogScore []PersonalBlogScore) error
}

// SetPersonalBlogScore 保存作业
func (Repo *Repository) SetPersonalBlogScore(personalBlogScore []PersonalBlogScore) error {
	result := Repo.DB.Create(&personalBlogScore)
	return result.Error
}




// GetPersonalBlogScoreByID 用ID获取个人博客成绩
func (Repo *Repository) GetPersonalBlogScoreByID(ID interface{}) (PersonalBlogScore, error) {
	var personalBlogScore PersonalBlogScore
	result := Repo.DB.Where("ID = ?", ID).Find(&personalBlogScore)
	return personalBlogScore, result.Error
}

// GetTeamBlogScoreByID 用ID获取团队博客成绩
func (Repo *Repository) GetTeamBlogScoreByID(ID interface{}) (TeamBlogScore, error) {
	var teamBlogScore TeamBlogScore
	result := Repo.DB.Where("ID = ?", ID).Find(&teamBlogScore)
	return teamBlogScore, result.Error
}

// GetPersonalBlogScores 获取全部个人博客成绩    仍需修改！
func (Repo *Repository) GetPersonalBlogScores(ID interface{}) (PersonalBlogScore, error) {
	var personalBlogScore PersonalBlogScore
	result := Repo.DB.Where("ID = ?", ID).Find(&personalBlogScore)
	return personalBlogScore, result.Error
}

// SetPersonalBlogScoreByID 根据ID设置个人博客成绩
func (Repo *Repository) SetPersonalBlogScoreByID(ID interface{}, grade int) (int64, error) {
	result := Repo.DB.Model(&PersonalBlogScore{}).Where("ID = ?", ID).Update("grade", grade)
	return result.RowsAffected, result.Error
}

// SetTeamBlogScoreByID 根据ID设置团队博客成绩
func (Repo *Repository) SetTeamBlogScoreByID(ID interface{}, grade int) (int64, error) {
	result := Repo.DB.Model(&TeamBlogScore{}).Where("ID = ?", ID).Update("grade", grade)
	return result.RowsAffected, result.Error
}
