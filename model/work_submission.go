//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
	"testing"
)

// WorkSubmission 作业提交模型
type WorkSubmission struct {
	gorm.Model
	SubmitterID  int    `gorm:"type:int;not null"`
	HomeworkID   int    `gorm:"type:int;not null"`
	SubmitStatus uint8  `gorm:"type:int;not null"`
	Filepath     string `gorm:"type:varchar(255)"`
}

// CreateWorkSubmission 创建作业提交表
func (Repo *Repository) CreateWorkSubmission(SubmitterID int, HomeworkID int,
	SubmitStatus uint8, Filepath string) (int64, error) {
	workSubmission := WorkSubmission{SubmitterID: SubmitterID, HomeworkID: HomeworkID,
		SubmitStatus: SubmitStatus, Filepath: Filepath}
	result := Repo.DB.Create(&workSubmission)
	return result.RowsAffected, result.Error
}

// GetWorkSubmissionByID 根据ID获取作业提交表
func (Repo *Repository) GetWorkSubmissionByID(ID int) (WorkSubmission, error) {
	var workSubmission WorkSubmission
	result := Repo.DB.Where("ID = ?", ID).Find(&workSubmission)
	return workSubmission, result.Error
}

// GetWorkSubmissionBySubmitterIDandHomeworkID 根据SubmitterID和HomeworkID获取作业提交表
func (Repo *Repository) GetWorkSubmissionBySubmitterIDandHomeworkID(SubmitterID int,
	HomeworkID int) (WorkSubmission, error) {
	var workSubmission WorkSubmission
	result := Repo.DB.Where(map[string]interface{}{"SubmitterID": SubmitterID,
		"HomeworkID": HomeworkID}).Find(&workSubmission)
	return workSubmission, result.Error
}

// SetSubmitStatusByID 根据ID设置作业提交状态
func (Repo *Repository) SetSubmitStatusByID(ID, submitStatus int) (*testing.T, error) {
	result := Repo.DB.Model(&WorkSubmission{}).Where("ID = ?", ID).Update("submit_status", submitStatus)
	return result.RowsAffected, result.Error
}

// DeleteWorkSubmissionByID 根据ID删除作业提交表
func (Repo *Repository) DeleteWorkSubmissionByID(ID int) error {
	result := Repo.DB.Delete(&WorkSubmission{}, ID)
	return result.Error
}
