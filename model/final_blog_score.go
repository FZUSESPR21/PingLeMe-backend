package model

import "gorm.io/gorm"

// FinalBlogScore 博客最终成绩模型
type FinalBlogScore struct {
	gorm.Model
	HomeworkID    uint `gorm:"type:int;not null"`
	ScorekeeperID uint `gorm:"type:int;not null"`
	BlogGrade     int  `gorm:"type:int;not null"`
}

// FinalScore 一次作业的最终成绩模型
type FinalScore struct {
	gorm.Model
}

type FinalBlogScoreRepositoryInterface interface {
	CountPersonalFinalScore(homework Homework, keeperID uint) (error, error)
	GetFinalBlogScore(homeworkID uint, keeperID uint) (FinalBlogScore, error)
}

// CountPersonalFinalScore 计算个人博客最终成绩(获取作业第一级并计算)
func (Repo *Repository) CountPersonalFinalScore(homework Homework, keeperID uint) (error, error, error) {
	var scoringItemsID []uint
	for _, item := range homework.ScoringItems {
		if item.Level == 1 {
			scoringItemsID = append(scoringItemsID, item.ID)
		}
	}
	var personalBlogScores []PersonalBlogScore
	result := Repo.DB.Where("scorekeeper_id = ? and scoring_item_id in ?", keeperID, scoringItemsID).Find(&personalBlogScores)
	var finalBlogScore FinalBlogScore
	for _, item := range personalBlogScores {
		finalBlogScore.BlogGrade += item.Grade
	}
	finalBlogScore.HomeworkID = homework.ID
	finalBlogScore.ScorekeeperID = keeperID
	deleteResult := Repo.DB.Where("homework_id = ? and scorekeeper_id = ? ", homework.ID, keeperID).Delete(FinalBlogScore{})
	finalResult := Repo.DB.Create(&finalBlogScore)
	return result.Error,deleteResult.Error, finalResult.Error
}

// GetFinalBlogScore 获取博客最终成绩
func (Repo *Repository) GetFinalBlogScore(homeworkID uint, keeperID uint) (FinalBlogScore, error) {
	var finalBlogScore FinalBlogScore
	result := Repo.DB.Where("homework_id = ? and scorekeeper_id = ?", homeworkID, keeperID).First(&finalBlogScore)
	return finalBlogScore, result.Error
}

// GetFinalScoreList 获取最终成绩列表（最终成绩包括评审表、博客、代码程序）
func (Repo *Repository) GetFinalScoreList() {

}
