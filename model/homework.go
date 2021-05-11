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
	HomeworkID  uint   `gorm:"type:int;not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Score       int    `gorm:"type:int;not null;default:-1"`
	Option      uint8  `gorm:"type:int;not null"`
	Note        string `gorm:"type:varchar(255)"`
	AssistantID uint   `gorm:"type:int;not null"`
	Level       int    `gorm:"not null;default:0"`
	Index       int    `gorm:"type:int;not null"`
}

type HomeworkRepositoryInterface interface {
	GetHomeworkByID(ID uint) (Homework, error)
	SetHomework(homework Homework) error
	GetAllHomeworkByPage(classID uint, page int, pageSize int) ([]Homework, error)
	CountHomework(classID uint) int
}

// GetHomeworkByID 获得某个特定ID的作业
func (Repo *Repository) GetHomeworkByID(ID uint) (Homework, error) {
	var homework Homework
	result := Repo.DB.First(&homework, ID)
	if result.Error != nil {
		return Homework{}, result.Error
	}

	var items []ScoringItem
	result = Repo.DB.Order("index desc").Where("homework_id = ?", ID).Find(&items)
	if result.Error != nil {
		return Homework{}, result.Error
	}
	homework.ScoringItems = items

	return homework, nil
}

// SetHomework 保存作业
func (Repo *Repository) SetHomework(homework Homework) error {
	result := Repo.DB.Create(&homework)
	return result.Error
}

// AssignedToAssistant 分配给助教
func (scoringItem *ScoringItem) AssignedToAssistant(assistantID uint) (ScoringItem, error) {
	scoringItem.AssistantID = assistantID
	return *scoringItem, nil
}

// GetAllHomeworkByPage 获得某个班级布置的所有作业中的一页
func (Repo *Repository) GetAllHomeworkByPage(classID uint, page int, pageSize int) ([]Homework, error) {
	var homework []Homework
	result := Repo.DB.Limit(pageSize).Offset((page-1)*pageSize).Where("class_id = ?", classID).Find(&homework)
	return homework, result.Error
}

// CountHomework 获得某个班级作业总数
func (Repo *Repository) CountHomework(classID uint) int {
	var sum int
	Repo.DB.Raw("select count(id) from homework where class_id = ? ", classID).Scan(&sum)
	return sum
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

//// GetSonScoringItems 获得某个评分项的所有下层子项
//func (scoringItem *ScoringItem) GetSonScoringItems() ([]ScoringItem, error) {
//	var scoringItems []ScoringItem
//	result := Repo.DB.Where("parent_item_id = ?", scoringItem.ID).Find(scoringItem)
//	return scoringItems, result.Error
//}

// AddHomework 布置新作业
func (Repo *Repository) AddHomework(homework Homework) error {
	result := Repo.DB.Create(&homework)
	return result.Error
}

//// AddScoringItem 增加评分项
//func (homework *Homework) AddScoringItem(scoringItem ScoringItem) error {
//	result := Repo.DB.Create(&scoringItem)
//	return result.Error
//}

// DeleteHomework 删除某个作业
func (Repo *Repository) DeleteHomework(homeworkID uint) error {
	result := Repo.DB.Delete(&Homework{}, homeworkID)
	return result.Error
}

//// DeleteScoringItem 删除该作业的某个评分项
//func (homework *Homework) DeleteScoringItem(scoringItemID uint) error {
//	result := Repo.DB.Delete(&ScoringItem{}, scoringItemID)
//	return result.Error
//}

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
