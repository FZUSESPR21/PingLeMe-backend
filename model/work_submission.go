//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// WorkSubmission 作业提交模型
type WorkSubmission struct {
	gorm.Model
	SubmitterID  uint   `gorm:"type:int;not null"`
	HomeworkID   uint   `gorm:"type:int;not null"`
	SubmitStatus uint8  `gorm:"type:int;not null"`
	Filepath     string `gorm:"type:varchar(255)"`
	IsReviewed   bool   `gorm:"default:false"`
}

type WorkSubmissionRepositoryInterface interface {
	CreateWorkSubmission(SubmitterID uint, HomeworkID uint, SubmitStatus uint8, Filepath string) (WorkSubmission, error)
	GetWorkSubmissionByID(ID interface{}) (WorkSubmission, error)
	GetWorkSubmissionBySubmitterIDandHomeworkID(SubmitterID int, HomeworkID int) (WorkSubmission, error)
	SetSubmitStatusByID(ID interface{}, submitStatus int) (int64, error)
	DeleteWorkSubmissionByID(ID interface{}) (int64, error)
}

// CreateWorkSubmission 创建作业提交表
func (Repo *Repository) CreateWorkSubmission(SubmitterID uint, HomeworkID uint,
	SubmitStatus uint8, Filepath string) (WorkSubmission, error) {
	workSubmission := WorkSubmission{SubmitterID: SubmitterID, HomeworkID: HomeworkID,
		SubmitStatus: SubmitStatus, Filepath: Filepath}
	result := Repo.DB.Create(&workSubmission)
	return workSubmission, result.Error
}

// GetWorkSubmissionByID 根据ID获取作业提交表
func (Repo *Repository) GetWorkSubmissionByID(ID interface{}) (WorkSubmission, error) {
	var workSubmission WorkSubmission
	result := Repo.DB.Where("ID = ?", ID).Find(&workSubmission)
	return workSubmission, result.Error
}

// GetWorkSubmissionBySubmitterIDandHomeworkID 根据SubmitterID和HomeworkID获取作业提交表
func (Repo *Repository) GetWorkSubmissionBySubmitterIDandHomeworkID(SubmitterID int,
	HomeworkID int) (WorkSubmission, error) {
	var workSubmission WorkSubmission
	result := Repo.DB.Where(map[string]interface{}{"SubmitterID": SubmitterID,
		"HomeworkID": HomeworkID}).First(&workSubmission)
	return workSubmission, result.Error
}

// GetWorkSubmissionsByHomeworkIDandClassID 根据HomeworkID和ClassID获取作业提交表列表
func (Repo *Repository) GetWorkSubmissionsByHomeworkIDandClassID(homeworkID int, ClassID int) ([]WorkSubmission, error) {
	var workSubmissions []WorkSubmission
	result := Repo.DB.Where("HomeworkID = ?", homeworkID).
		Where("SubmitterID IN ?", Repo.DB.Table("student_class").Select("user_id").Where("class_id = ?", ClassID)).
		Find(&workSubmissions)
	return workSubmissions, result.Error
}

// SetSubmitStatusByID 根据ID设置作业提交状态
func (Repo *Repository) SetSubmitStatusByID(ID interface{}, submitStatus int) (int64, error) {
	result := Repo.DB.Model(&WorkSubmission{}).Where("ID = ?", ID).Update("submit_status", submitStatus)
	return result.RowsAffected, result.Error
}

// DeleteWorkSubmissionByID 根据ID删除作业提交表
func (Repo *Repository) DeleteWorkSubmissionByID(ID interface{}) (int64, error) {
	result := Repo.DB.Delete(&WorkSubmission{}, ID)
	return result.RowsAffected, result.Error
}
