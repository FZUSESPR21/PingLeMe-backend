//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// PersonalBlogScore 个人博客成绩模型
type PersonalBlogScore struct {
	gorm.Model
	ScoringItemID uint `gorm:"type:int;not null"`
	ScorekeeperID uint `gorm:"type:int;not null"`
	Grade         float32  `gorm:"type:float;not null"`
}

// TeamBlogScore 团队博客成绩模型
type TeamBlogScore struct {
	gorm.Model
	ScoringItemID uint `gorm:"type:int;not null"`
	ScorekeeperID uint `gorm:"type:int;not null"`
	Grade         float32  `gorm:"type:float;not null"`
}

type BlogScoreRepositoryInterface interface {
	SetPersonalBlogScore(personalBlogScore []PersonalBlogScore) error
	SetTeamBlogScore(teamBlogScore []TeamBlogScore) error
	LoadPersonalZeroScore(homeworkID uint, scorekeeperID uint) (error,error,error)
	LoadTeamZeroScore(homeworkID uint, scorekeeperID uint) (error,error,error)
}

// LoadPersonalZeroScore 批改某人的一份作业前，若无记录，则先载入该人该作业所有评分项成绩(成绩初始设置为0)
func (Repo *Repository) LoadPersonalZeroScore(homeworkID uint, scorekeeperID uint) (error,error,error) {
	var homework Homework
	result1 := Repo.DB.Where("id = ?", homeworkID).First(&homework)
	var personalBlogScore PersonalBlogScore
	result2 := Repo.DB.Where("scoring_item_id = ? and scorekeeper_id = ?", homework.ScoringItems[0].ID, scorekeeperID).First(&personalBlogScore)
	var zeroScoreItems []PersonalBlogScore
	if personalBlogScore.ScorekeeperID != scorekeeperID {
		for _, item := range homework.ScoringItems {
			var aim PersonalBlogScore
			aim.ScorekeeperID = scorekeeperID
			aim.ScoringItemID = item.ID
			aim.Grade = 0
			zeroScoreItems = append(zeroScoreItems, aim)
		}
	}
	result3 := Repo.DB.Create(&zeroScoreItems)
	return result1.Error, result2.Error, result3.Error
}

// LoadTeamZeroScore 批改某团队的一份作业前，若无记录，则先载入该人该作业所有评分项成绩(成绩初始设置为0)
func (Repo *Repository) LoadTeamZeroScore(homeworkID uint, scorekeeperID uint) (error,error,error) {
	var homework Homework
	result1 := Repo.DB.Where("id = ?", homeworkID).First(&homework)
	var teamBlogScore TeamBlogScore
	result2 := Repo.DB.Where("scoring_item_id = ? and scorekeeper_id = ?", homework.ScoringItems[0].ID, scorekeeperID).First(&teamBlogScore)
	var zeroScoreItems []TeamBlogScore
	if teamBlogScore.ScorekeeperID != scorekeeperID {
		for _, item := range homework.ScoringItems {
			var aim TeamBlogScore
			aim.ScorekeeperID = scorekeeperID
			aim.ScoringItemID = item.ID
			aim.Grade = 0
			zeroScoreItems = append(zeroScoreItems, aim)
		}
	}
	result3 := Repo.DB.Create(&zeroScoreItems)
	return result1.Error, result2.Error, result3.Error
}

// SetPersonalBlogScore 保存个人作业成绩（接收前端发来的评分结果项）(useless)
func (Repo *Repository) SetPersonalBlogScore(personalBlogScore []PersonalBlogScore) error {
	result := Repo.DB.Create(&personalBlogScore)
	return result.Error
}

// SetTeamBlogScore 保存团队作业成绩
func (Repo *Repository) SetTeamBlogScore(teamBlogScore []TeamBlogScore) error {
	result := Repo.DB.Create(&teamBlogScore)
	return result.Error
}


// GetPersonalBlogScoreByID 用ID获取个人博客成绩
func (Repo *Repository) GetPersonalBlogScoreByID(ID interface{}) (PersonalBlogScore, error) {
	var personalBlogScore PersonalBlogScore
	result := Repo.DB.Where("id = ?", ID).Find(&personalBlogScore)
	return personalBlogScore, result.Error
}

// GetTeamBlogScoreByID 用ID获取团队博客成绩
func (Repo *Repository) GetTeamBlogScoreByID(ID interface{}) (TeamBlogScore, error) {
	var teamBlogScore TeamBlogScore
	result := Repo.DB.Where("id = ?", ID).Find(&teamBlogScore)
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
