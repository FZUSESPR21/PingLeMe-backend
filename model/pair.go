//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"errors"
	"gorm.io/gorm"
)

// Pair 结对模型
type Pair struct {
	gorm.Model
	Student1ID uint `gorm:"type:int;not null;index:studentID"`
	Student2ID uint `gorm:"type:int;index:studentID"`
}

type PairRepositoryInterface interface {
	GetPair(ID interface{}) (Pair, error)
	CreatePair(pair Pair) (Pair, error)
	GetPairByStudentID(ID uint) (uint, error)
	DeletePair(ID int) error
	DeletePairByStudentID(ID int) error
	UpdatePair(ID int, student1ID uint, student2ID uint) (Pair, error)
	UpdatePairByStu(student1ID uint, student2ID uint) (int, error)
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
func (Repo *Repository) GetPairByStudentID(ID uint) (uint, error) {
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
func (Repo *Repository) UpdatePair(ID int, student1ID uint, student2ID uint) (Pair, error) {
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

// UpdatePairByStu 更新Pair
func (Repo *Repository) UpdatePairByStu(student1ID uint, student2ID uint) (int, error) {
	//1为成功
	//2为对方已与别人结对
	//3为保存修改错误
	//4为添加结对失败

	var pair1 Pair
	s1 := 1
	result1 := Repo.DB.Where("student1_id = ?", student1ID).Or("student2_id = ?", student1ID).First(&pair1)
	if result1.Error != nil {
		s1 = 0
	}

	var pair2 Pair
	s2 := 1
	result2 := Repo.DB.Where("student1_id = ?", student2ID).Or("student2_id = ?", student2ID).First(&pair2)
	if result2.Error != nil {
		s2 = 0
	}


	if s2 == 1 {
		return 2, nil //对方已和别人结对
	} else if s1 == 1 {
		if pair1.Student1ID == student1ID {
			pair1.Student2ID = student2ID
		} else if pair1.Student2ID == student1ID {
			pair1.Student1ID = student2ID
			result1 = Repo.DB.Save(&pair1)
			if result1.Error != nil {
				return 3, result1.Error//保存修改错误
			}
		}//修改队友信息成功
	} else {
		var pair Pair
		pair.Student1ID = student1ID
		pair.Student2ID = student2ID
		result := Repo.DB.Create(&pair)//两个都未结对，添加结对
		if result.Error != nil {
			return 4, result.Error//添加结对失败
		}
	}
	return 1, nil//成功
}
