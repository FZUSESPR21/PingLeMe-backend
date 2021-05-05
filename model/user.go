//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UID            string `gorm:"type:varchar(9);not null;unique"`
	PasswordDigest string `gorm:"type:varchar(30);not null"`
	Nickname       string `gorm:"type:varchar(20);not null;unique"`
	Role           uint8  `gorm:"type:int;default:0;not null"`
	Roles          []Role `gorm:"many2many:user_role"`
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

// SetUsers 添加用户组
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


func (Repo *Repository) GetAllTeacher() []User {
	var user []User
	result := Repo.DB.Where("uid = ?", UID).First(&user)
	return user
}