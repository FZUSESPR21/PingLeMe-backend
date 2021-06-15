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
	UserID  uint
	ClassID uint
}

// StudentClass 学生-班级
type StudentClass struct {
	UserID  uint
	ClassID uint
}

type AssistantInfo struct {
	AssistantID   uint
	AssistantName string
}

type ClassInfo struct {
	ClassID     uint            `json:"class_id"`
	ClassName   string          `json:"class_name"`
	TeacherID   uint            `json:"teacher_id"`
	TeacherName string          `json:"teacher_name"`
	Assistants  []AssistantInfo `json:"assistants"`
	PairStatus  bool            `json:"pair_status"`
	TeamStatus  bool            `json:"team_status"`
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
	EditStuClass(studentID int, newClassID int) error
	GetStusByClassName(classID int) (int, []User, error)
	GetAssisByClassName(classID int) (int, []User, error)
	GetTeacherByClassID(classID int) (User, error)
	GetClassInfoList() ([]ClassInfo, error)
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
		result := Repo.DB.Exec("INSERT IGNORE INTO student_class (class_id, student_id) values(?, ?)", stuClass.ClassID, stuClass.UserID)
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

// EditStuClass 修改学生班级
func (Repo *Repository) EditStuClass(studentID int, newClassID int) error {
	result := Repo.DB.Exec("update student_class set class_id = ? where user_id = ?", newClassID, studentID)
	return result.Error
}

// GetStusByClassName 查看班级学生列表
func (Repo *Repository) GetStusByClassName(classID int) (int, []User, error) {
	var stus []User
	var studentClass []StudentClass
	result := Repo.DB.Table("student_class").Where("class_id = ?", classID).Find(&studentClass)

	for i := 0; i < len(studentClass); i++ {
		var stu User
		result = Repo.DB.Where("id = ?", studentClass[i].UserID).First(&stu)
		stus = append(stus, stu)
	}
	return len(studentClass), stus, result.Error
}

// GetAssisByClassName 查看班级助教列表
func (Repo *Repository) GetAssisByClassName(classID int) (int, []User, error) {
	var assis []User
	var teacherClass []TeacherClass
	result := Repo.DB.Table("teacher_class").Where("class_id = ?", classID).Find(&teacherClass)

	var num int
	for i := 0; i < len(teacherClass); i++ {
		var stu User
		result = Repo.DB.Where("id = ? and role = 2", teacherClass[i].UserID).First(&stu)
		if result.RowsAffected > 0 {
			num++
			assis = append(assis, stu)
		}
	}
	if num > 0 {
		return num, assis, nil
	} else {
		return num, assis, result.Error
	}
}

//GetTeacherByClassID 通过ClassID查找老师
func (Repo *Repository) GetTeacherByClassID(classID int) (User, error) {
	var teacher User
	var teacherClass []TeacherClass
	result := Repo.DB.Table("teacher_class").Where("class_id = ?", classID).Find(&teacherClass)
	var num int
	for i := 0; i < len(teacherClass); i++ {
		result = Repo.DB.Where("id = ? and role = 1", teacherClass[i].UserID).First(&teacher)
		if result.RowsAffected > 0 {
			num++
		}
	}
	if num > 0 {
		return teacher, nil
	} else {
		return teacher, result.Error
	}

}
func (Repo *Repository) GetClassInfoList() ([]ClassInfo, error) {
	var classes []Class
	_ = Repo.DB.Find(&classes)

	classInfo := make([]ClassInfo, len(classes))
	for index, class := range classes {
		classInfo[index].ClassID = class.ID
		classInfo[index].ClassName = class.Name
		classInfo[index].Assistants = make([]AssistantInfo, 0)
		rows, err := Repo.DB.Raw("SELECT * FROM `users` WHERE id IN (SELECT user_id FROM `teacher_class` WHERE class_id = ?)", class.ID).Rows()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var user User
			err := Repo.DB.ScanRows(rows, &user)
			if err != nil {
				return nil, err
			}
			if user.Role == RoleTeacher {
				classInfo[index].TeacherName = user.UserName
				classInfo[index].TeacherID = user.ID
			} else {
				classInfo[index].Assistants = append(classInfo[index].Assistants, AssistantInfo{
					AssistantID:   user.ID,
					AssistantName: user.UserName,
				})
			}
		}
	}
	return classInfo, nil
}
