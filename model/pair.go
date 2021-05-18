//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"errors"
	"fmt"
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
	UpdatePairByStu(student1ID int, student2ID int) (int, error)
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

// UpdatePair 更新Pair
func (Repo *Repository) UpdatePairByStu(student1ID int, student2UID int) (int, error) {
	//0为查询错误
	//1为成功
	//2为对方已与别人结对
	//3为保存错误
	var user UserRepositoryInterface
	user1, err := user.GetUserByUID(fmt.Sprint(student2UID))
	if err != nil {
		return 0, err
	}
	student2IDu := user.GetUserID(user1)
	student2ID := int(student2IDu)

	var pair1 Pair
	result1 := Repo.DB.Where("student1_id = ?", student1ID).Or("student2_id = ?", student1ID).First(&pair1)
	if result1.Error != nil {
		return 0, result1.Error
	}

	var pair2 Pair
	result2 := Repo.DB.Where("student1_id = ?", student2ID).Or("student2_id = ?", student2ID).First(&pair2)
	if result2.Error != nil {
		return 0, result2.Error
	}
	var t int
	t = 0
	if pair2.Student1ID == student2ID && (pair2.Student2ID == 0 || pair2.Student2ID == student1ID) {
		t = 1
	}
	if (pair2.Student1ID == 0 || pair2.Student1ID == student1ID) && pair2.Student2ID == student2ID {
		t = 1
	}

	if pair1.Student1ID == student1ID && t == 1 {
		pair1.Student2ID = student2ID
	} else if pair1.Student2ID == student1ID && t == 1 {
		pair1.Student1ID = student2ID
	} else {
		return 2, nil //对方已和别人结对

	}
	result1 = Repo.DB.Save(&pair1)
	if result1.Error != nil {
		return 3, result1.Error
	}
	return 1, nil
}
