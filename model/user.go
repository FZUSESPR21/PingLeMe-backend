//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UID            string `gorm:"not null;unique"`
	PasswordDigest string `gorm:"not null"`
	UserName       string `gorm:"type:varchar(20);not null"`
	Role           uint8  `gorm:"type:int;default:0;not null"`
	Roles          []Role `gorm:"many2many:user_role"`
}

type Teacher struct {
	UID       string `json:"uid"`
	UserName  string `json:"user_name"`
	ClassID   uint   `json:"class_id"`
	ClassName string `json:"class_name"`
}

type Assistant struct {
	UID       string `json:"uid"`
	UserName  string `json:"user_name"`
	ClassID   uint   `json:"class_id"`
	ClassName string `json:"class_name"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12

	// RoleTeacher 身份：教师
	RoleTeacher = 1
	// RoleAssistant 身份：助教
	RoleAssistant = 2
	// RoleAdmin 身份：超级管理员
	RoleAdmin = 9
	// RoleStudent 身份：学生
	RoleStudent = 0
)

type UserRepositoryInterface interface {
	GetUser(ID interface{}) (User, error)
	GetUserByUID(UID string) (User, error)
	SetUser(user User) error
	SetUsers(user []User) error
	DeleteUser(ID interface{}) error
	GetAllTeacher() (int64, []Teacher, error)
	GetAllAssistant() (int64, []Assistant, error)
	AddTeacherByUser(teacher User) (int64, error)
	ChangeUserPassword(user User, newPasswordDigest string) error
	GetUserTeamID(user User) (uint, error)
	GetStudentClassID(userID uint) (uint, error)
}

// GetUser 用ID获取用户
func (Repo *Repository) GetUser(ID interface{}) (User, error) {
	var user User
	result := Repo.DB.First(&user, ID)
	return user, result.Error
}

// GetUserByUID 用UID获取用户
func (Repo *Repository) GetUserByUID(UID string) (User, error) {
	var user User
	result := Repo.DB.Where("uid = ?", UID).First(&user)
	return user, result.Error
}

// SetUser 添加用户
func (Repo *Repository) SetUser(user User) error {
	result := Repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	err := Repo.SetUserRole(user.Role, user)
	if err != nil {
		return err
	}
	return nil
}

// SetUsers 添加用户组,
func (Repo *Repository) SetUsers(users []User) error {
	result := Repo.DB.Create(&users)
	if result.Error != nil {
		return result.Error
	}
	err := Repo.SetUsersRole(users[0].Role, users)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (Repo *Repository) DeleteUser(ID interface{}) error {
	Repo.DB.Delete(&User{}, ID)
	return nil
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (Repo *Repository) GetAllTeacher() (int64, []Teacher, error) {
	var users []User
	var teachers []Teacher
	result := Repo.DB.Where("role = ?", RoleTeacher).Find(&users)
	for _, user := range users {
		row := Repo.DB.Raw("SELECT class_id FROM `teacher_class` WHERE user_id = ? LIMIT 1", user.ID).Row()
		var classID uint
		err := row.Scan(&classID)
		if err == nil {
			class, err := Repo.GetClassByID(classID)
			if err != nil {
				return 0, nil, err
			}
			teachers = append(teachers, Teacher{
				UID:       user.UID,
				UserName:  user.UserName,
				ClassID:   class.ID,
				ClassName: class.Name,
			})
		} else {
			teachers = append(teachers, Teacher{
				UID:       user.UID,
				UserName:  user.UserName,
				ClassID:   0,
				ClassName: "",
			})
		}
	}
	return result.RowsAffected, teachers, result.Error
}

func (Repo *Repository) GetAllAssistant() (int64, []Assistant, error) {
	var users []User
	var assistant []Assistant
	result := Repo.DB.Where("role = ?", RoleAssistant).Find(&users)
	for _, user := range users {
		row := Repo.DB.Raw("SELECT class_id FROM `teacher_class` WHERE user_id = ? LIMIT 1", user.ID).Row()
		var classID uint
		err := row.Scan(&classID)
		if err == nil {
			class, err := Repo.GetClassByID(classID)
			if err != nil {
				return 0, nil, err
			}
			assistant = append(assistant, Assistant{
				UID:       user.UID,
				UserName:  user.UserName,
				ClassID:   class.ID,
				ClassName: class.Name,
			})
		} else {
			assistant = append(assistant, Assistant{
				UID:       user.UID,
				UserName:  user.UserName,
				ClassID:   0,
				ClassName: "",
			})
		}

	}
	return result.RowsAffected, assistant, result.Error
}

func (Repo *Repository) AddTeacherByUser(teacher User) (int64, error) {
	result := Repo.DB.Create(&teacher)
	return result.RowsAffected, result.Error
}

// ChangeUserPassword 修改用户密码
func (Repo *Repository) ChangeUserPassword(user User, newPasswordDigest string) error {
	result := Repo.DB.Model(&user).Update("password_digest", newPasswordDigest)
	return result.Error
}
func (Repo *Repository) GetUserTeamID(user User) (uint, error) {
	var teamID uint
	row := Repo.DB.Table("student_team").Where("student_id = ?", user.ID).Select("team_id").Row()
	err := row.Scan(&teamID)
	if err != nil {
		return 0, err
	} else {
		return teamID, nil
	}
}

func (Repo *Repository) GetStudentClassID(userID uint) (uint, error) {
	raw := Repo.DB.Raw("SELECT `class_id` FROM `student_class` WHERE `user_id` = ? LIMIT 1", userID).Row()
	classID := uint(0)
	err := raw.Scan(&classID)
	return classID, err
}
