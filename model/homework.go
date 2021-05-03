//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
	"time"
)

// Homework 作业模型
type Homework struct {
	gorm.Model
	ClassID      uint      `gorm:"type:int;not null"`
	Type         uint8     `gorm:"type:int;not null"`
	Title        string    `gorm:"type:varchar(255);not null"`
	Content      string    `gorm:"type:text;not null"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
	ScoringItems []ScoringItem
}

// ScoringItem 评分项模型
type ScoringItem struct {
	gorm.Model
	HomeworkID   uint   `gorm:"type:int;not null"`
	Description  string `gorm:"type:varchar(255);not null"`
	Score        int    `gorm:"type:int;not null"`
	Option       uint8  `gorm:"type:int;not null"`
	Note         string `gorm:"type:varchar(255)"`
	AssistantID  uint   `gorm:"type:int;not null"`
	ParentItemID uint   `gorm:"type:int;not null"`
	Sequence     int    `gorm:"type:int;not null"`
}

// GetHomeworkByID 获得某个特定ID的作业
func (Repo *Repository) GetHomeworkByID(ID interface{}) (Homework, error) {
	var homework Homework
	result := Repo.DB.First(&homework, ID)
	return homework, result.Error
}

// GetAllHomework 获得某个班级布置的所有作业
func (class *Class) GetAllHomework() ([]Homework, error) {
	var homework []Homework
	result := Repo.DB.Where("class_id = ?", class.ID).Find(&homework)
	return homework, result.Error
}

// GetAllScoringItem 获得某个作业的所有评分项
func (homework *Homework) GetAllScoringItem() ([]ScoringItem, error) {
	//var scoringItem []ScoringItem
	//result := Repo.DB.Where("homework_id = ?", homework.ID).Find(scoringItem)
	return homework.ScoringItems, nil
}

// GetAssignedScoringItem 获得分配给某个助教所有的评分项
func (user *User) GetAssignedScoringItem() ([]ScoringItem, error) {
	var scoringItem []ScoringItem
	result := Repo.DB.Where("assistant_id = ?", user.ID).Find(scoringItem)
	return scoringItem, result.Error
}

// GetSonScoringItems 获得某个评分项的所有下层子项
func (scoringItem *ScoringItem) GetSonScoringItems() ([]ScoringItem, error) {
	var scoringItems []ScoringItem
	result := Repo.DB.Where("parent_item_id = ?", scoringItem.ID).Find(scoringItem)
	return scoringItems, result.Error
}

// AddHomework 布置新作业
func (Repo *Repository) AddHomework(homework Homework) error {
	result := Repo.DB.Create(&homework)
	return result.Error
}

// AddScoringItem 增加评分项
func (homework *Homework) AddScoringItem(scoringItem ScoringItem) error {
	result := Repo.DB.Create(&scoringItem)
	return result.Error
}

// DeleteHomework 删除某个作业
func (Repo *Repository) DeleteHomework(homeworkID uint) error {
	result := Repo.DB.Delete(&Homework{}, homeworkID)
	return result.Error
}

// DeleteScoringItem 删除该作业的某个评分项
func (homework *Homework) DeleteScoringItem(scoringItemID uint) error {
	result := Repo.DB.Delete(&ScoringItem{}, scoringItemID)
	return result.Error
}

// UpdateHomework 更改作业信息
func (Repo *Repository) UpdateHomework(homework Homework) error {
	result := Repo.DB.Model(&homework).Updates(homework)
	return result.Error
}

// UpdateScoringItem 更改评分项信息
func (Repo *Repository) UpdateScoringItem(scoringItem ScoringItem) error {
	result := Repo.DB.Model(&scoringItem).Updates(scoringItem)
	return result.Error
}
