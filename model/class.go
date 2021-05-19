//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"PingLeMe-Backend/util"
	"gorm.io/gorm"
)

// Class 班级模型
type Class struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;unique"`
	Teachers []User `gorm:"many2many:teacher_class;"`
	Students []User `gorm:"many2many:student_class;"`
}

// TeacherClass 教师-班级
type TeacherClass struct {
	TeacherID uint
	ClassID   uint
}

// StudentClass 学生-班级
type StudentClass struct {
	StudentID uint
	ClassID   uint
}

type ClassRepositoryInterface interface {
	GetClassByID(ID interface{}) (Class, error)
	AddClass(name string) (Class, error)
	DeleteClass(classID interface{}) error
	UpdateClassName(class Class, name string) error
	GetClassByName(name string) (Class, error)
	AddStudent(class Class, student User) error
	AddTeacher(class Class, teacher User) error
	DeleteTeacher(class Class, teacher User) error
}

// GetClassByID 通过班级ID获取班级
func (Repo *Repository) GetClassByID(ID interface{}) (Class, error) {
	var class Class
	result := Repo.DB.First(&class, ID)
	return class, result.Error
}

// GetClassByName 通过班级名称获取班级
func (Repo *Repository) GetClassByName(name string) (Class, error) {
	var class Class
	result := Repo.DB.Where("Name = ?", name).First(&class)
	return class, result.Error
}

// AddClass 添加一个班级
func (Repo *Repository) AddClass(name string) (Class, error) {
	class := Class{Name: name}
	result := Repo.DB.Create(&class)
	return class, result.Error
}

// ClassAddStudents 班级批量添加学生
func (Repo *Repository) ClassAddStudents(stuClasses []StudentClass) []error {
	errs := make([]error, 0)
	for index, stuClass := range stuClasses {
		result := Repo.DB.Exec("INSERT IGNORE INTO student_class (class_id, student_id) values(?, ?)", stuClass.ClassID, stuClass.StudentID)
		if result.Error != nil {
			errs = append(errs, result.Error)
		} else if result.RowsAffected == 0 {
			errs = append(errs, &util.RecordAlreadyExistErr{Row: index})
		}
	}
	return errs
}

// AddTeacher 添加一个老师
func (Repo *Repository) AddTeacher(class Class, teacher User) error {
	var classID = class.ID
	var teacherID = teacher.ID
	result := Repo.DB.Exec("insert into teacher_class(class_id,user_id) values(?,?)", classID, teacherID)
	return result.Error
}

// AddStudent 添加一个学生
func (Repo *Repository) AddStudent(class Class, student User) error {
	var classID = class.ID
	var studentID = student.ID
	result := Repo.DB.Exec("insert into student_class(class_id,user_id) values(?,?)", classID, studentID)
	return result.Error
}

// DeleteClass 删除班级
func (Repo *Repository) DeleteClass(classID interface{}) error {
	result := Repo.DB.Delete(&Class{}, classID)
	return result.Error
}

// DeleteTeacher 删除该班级里的一个老师
func (Repo *Repository) DeleteTeacher(class Class, teacher User) error {
	var classID = class.ID
	var teacherID = teacher.ID
	result := Repo.DB.Exec("delete from teacher_class where class_id = ? and user_id = ?", classID, teacherID)
	return result.Error
}

// DeleteStudent 删除改班级里的一个学生
func (Repo *Repository) DeleteStudent(class Class, student User) error {
	var classID = class.ID
	var studentID = student.ID
	result := Repo.DB.Exec("delete from student_class where class_id = ? and student_id = ?", classID, studentID)
	return result.Error
}

// UpdateClassName 修改班级名字
func (Repo *Repository) UpdateClassName(class Class, name string) error {
	result := Repo.DB.Model(&class).Update("name", name)
	return result.Error
}
