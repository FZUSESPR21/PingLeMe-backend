//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"PingLeMe-Backend/util"
	"errors"
	"reflect"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Role 角色模型
type Role struct {
	gorm.Model
	Type       uint        `gorm:"type:int;not null;unique"`
	Permission []Permission `gorm:"many2many:role_permission;"`
	User       []User       `gorm:"many2many:user_role;"`
	Desc       string       `gorm:"unique;"`
}

type UserRole struct {
	RoleID		uint
	UserID		uint
}

// Permission 权限模型
type Permission struct {
	gorm.Model
	Desc string `gorm:"unique;"`
	Type uint   `gorm:"type:int;not null;unique"`
	Role []Role `gorm:"many2many:role_permission;"`
}

type RBACRepositoryInterface interface {
	GetUserRoles(ID uint) ([]Role, error)
	GetUserPermissions(ID uint) ([]Permission, error)
}

// SetRole 新增角色
func (Repo *Repository) SetRole(roleType uint, roleDesc string) error {
	result := Repo.DB.Create(Role{
		Type: roleType,
		Desc: roleDesc,
	})

	if result.Error != nil {
		util.Log().Error("rbac.go/SetRole", zap.Error(result.Error))
	}

	return result.Error
}

// SetPermission 新增权限
func (Repo *Repository) SetPermission(permissionType uint, permissionDesc string) error {
	result := Repo.DB.Create(Permission{
		Type: permissionType,
		Desc: permissionDesc,
	})

	if result.Error != nil {
		util.Log().Error("rbac.go/SetPermission", zap.Error(result.Error))
	}

	return result.Error
}

// SetUserRole 设置用户角色
func (Repo *Repository) SetUserRole(roleType uint8, user User) error {
	var role Role
	result := Repo.DB.Where("type = ?", roleType).First(&role)
	if result.Error != nil {
		util.Log().Error(result.Error.Error())
		return result.Error
	}
	result = nil
	var userRole UserRole
	result = Repo.DB.Table("user_role").Where("role_id = ? AND user_id = ?", role.ID, user.ID).First(&userRole)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		Repo.DB.Exec("INSERT INTO user_role (role_id, user_id) VALUES (?, ?)", role.ID, user.ID)
	}

	return nil
}

// SetUsersRole 设置用户角色
func (Repo *Repository) SetUsersRole(roleType uint8, users []User) error {
	var role Role
	result := Repo.DB.Where("type = ?", roleType).First(&role)
	if result.Error != nil {
		util.Log().Error("rbac.go/SetUserRole", zap.Error(result.Error))
		return result.Error
	}

	for _, user := range users {
		Repo.DB.Exec("INSERT IGNORE INTO user_role (role_id, user_id) VALUES (?, ?)", role.ID, user.ID)
	}

	return nil
}

// GetUserRoles 获得用户角色
func (Repo *Repository) GetUserRoles(ID uint) ([]Role, error) {
	var user User
	result := Repo.DB.Preload("Roles").Where("id = ?", ID).Find(&user)
	return user.Roles, result.Error
}

// SetRolePermissions 设置角色权限
func (Repo *Repository) SetRolePermissions(roleDescOrType interface{}, permissions []Permission) error {
	var rolePermission Role
	switch desOrType := roleDescOrType.(type) {
	case uint:
		rolePermission = Role{
			Type:       desOrType,
			Permission: permissions,
		}
	case string:
		rolePermission = Role{
			Desc:       desOrType,
			Permission: permissions,
		}
	default:
		return &util.InterfaceTypeErr{Name: reflect.TypeOf(roleDescOrType).String()}
	}

	result := Repo.DB.Create(rolePermission)

	if result.Error != nil {
		util.Log().Error("rbac.go/SetRolePermission", zap.Error(result.Error))
	}

	return result.Error
}

// GetRolePermissions 获取角色权限
func (Repo *Repository) GetRolePermissions(roleDescOrType interface{}) ([]Permission, error) {
	var role Role
	switch roleDescOrType.(type) {
	case uint8:
		Repo.DB.Preload("Permission").Where("type = ?", roleDescOrType).First(&role)
	case string:
		Repo.DB.Preload("Permission").Where("desc = ?", roleDescOrType).First(&role)
	default:
		return nil, &util.InterfaceTypeErr{Name: reflect.TypeOf(roleDescOrType).String()}
	}

	return role.Permission, nil
}

// GetUserPermissions 获取用户权限
func (Repo *Repository) GetUserPermissions(ID uint) ([]Permission, error) {
	roles, err := Repo.GetUserRoles(ID)
	if err != nil {
		util.Log().Error("model/rbac.go/GetUserPermissions", zap.Error(err))
		return nil, err
	}
	permissions := make([]Permission, 0)
	for _, role := range roles {
		permission, err := Repo.GetRolePermissions(role.Type)
		if err != nil {
			util.Log().Error("model/rbac.go/GetUserPermissions", zap.Error(err))
			return nil, err
		}
		permissions = append(permissions, permission...)
	}
	return permissions, nil
}
