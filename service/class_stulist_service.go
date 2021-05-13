//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// ClassStuList 班级学生列表
type ClassStuList struct {
	model.ClassRepositoryInterface
	ClassName string `json:"className"`
}

// StuListOfClass 班级学生列表
func (service *ClassStuList) StuListOfClass() ([]model.User, error) {
	stus,err := service.GetStusByClassName(service.ClassName)
	if err != nil {
		return nil,err
	}
	return stus, nil
}