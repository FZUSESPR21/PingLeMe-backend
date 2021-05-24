//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// ClassStuList 班级学生列表
type ClassStuList struct {
	model.ClassRepositoryInterface
}

// StuListOfClass 班级学生列表
func (service *ClassStuList) StuListOfClass(classID int) ([]model.User, error) {
	stus, err := service.GetStusByClassName(classID)
	if err != nil {
		return nil, err
	}
	return stus, nil
}
