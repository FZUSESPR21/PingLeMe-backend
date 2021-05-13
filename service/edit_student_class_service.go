//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
)

// EditStudentClassService 改变学生班级
type EditStudentClassService struct {
	model.ClassRepositoryInterface
	StudentID int `json:"uid"`
	NewClass int `json:"newClass"`
}

// EditStudentClass 改变学生班级
func (service *EditStudentClassService) EditStudentClass() error {
	err := service.EditStuClass(service.StudentID, service.NewClass)
	if err != nil {
		return err
	}
	return nil
}
