//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"errors"
	"gorm.io/gorm"
)

// Pair 结对模型
type Pair struct {
	gorm.Model
	Student1ID int `gorm:"type:int;not null;index:studentID"`
	Student2ID int `gorm:"type:int;index:studentID"`
}

type PairRepositoryInterface interface {
	GetPair(ID interface{}) (Pair, error)
	CreatePair(pair Pair) (Pair, error)
	GetPairByStudentID(ID int) (int, error)
	DeletePair(ID int) error
	DeletePairByStudentID(ID int) error
	UpdatePair(ID int, student1ID int, student2ID int) (Pair, error)
}

// GetPair 用ID获取结对
func (Repo *Repository) GetPair(ID interface{}) (Pair, error) {
	var pair Pair
	result := Repo.DB.First(&pair, ID)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}

// CreatePair 创建结对
func (Repo *Repository) CreatePair(pair Pair) (Pair, error) {
	result := Repo.DB.Where("student1_id = ?", pair.Student1ID).Or("student2_id = ?", pair.Student1ID).First(&pair)
	if result.RowsAffected != 0 {
		return pair, errors.New("该学生已结对")
	}
	result = Repo.DB.Where("student1_id = ?", pair.Student2ID).Or("student2_id = ?", pair.Student2ID).First(&pair)
	if result.RowsAffected != 0 {
		return pair, errors.New("该学生已结对")
	}

	result = Repo.DB.Create(&pair)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}

// GetPairByStudentID 根据学生ID获取结对
func (Repo *Repository) GetPairByStudentID(ID int) (int, error) {
	var pair Pair
	result := Repo.DB.Where("student1_id = ?", ID).Or("student2_id = ?", ID).First(&pair)
	if result.Error != nil {
		return 0, result.Error
	}
	if pair.Student1ID == ID {
		return pair.Student2ID, nil
	} else {
		return pair.Student1ID, nil
	}
}

// DeletePair 根据ID删除结对
func (Repo *Repository) DeletePair(ID int) error {
	result := Repo.DB.Delete(&Pair{}, ID)
	return result.Error
}

// DeletePairByStudentID 根据学生ID删除结对
func (Repo *Repository) DeletePairByStudentID(ID int) error {
	result := Repo.DB.Where("student1_id = ?", ID).Or("student2_id = ?", ID).Delete(&Pair{})
	return result.Error
}

// UpdatePair 更新Pair
func (Repo *Repository) UpdatePair(ID int, student1ID int, student2ID int) (Pair, error) {
	var pair Pair
	result := Repo.DB.First(&pair, ID)
	if result.Error != nil {
		return Pair{}, result.Error
	}

	if student1ID != 0 {
		pair.Student1ID = student1ID
	}
	if student2ID != 0 {
		pair.Student2ID = student2ID
	}
	result = Repo.DB.Save(&pair)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}
