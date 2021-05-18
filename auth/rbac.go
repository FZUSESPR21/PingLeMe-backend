//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package auth

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/util"
	"reflect"
)

type RBACAuth struct {
	model.RBACRepositoryInterface
}

// CheckUserRole 检查用户角色
func (rbac RBACAuth) CheckUserRole(user model.User, roleDescOrType interface{}) (bool, error) {
	roles, err := rbac.GetUserRoles(user.ID)
	if err != nil {
		return false, err
	}

	switch roleDescOrType.(type) {
	case uint8:
		for _, r := range roles {
			if r.Type == roleDescOrType {
				return true, nil
			}
		}
		return false, nil
	case string:
		for _, r := range roles {
			if r.Desc == roleDescOrType {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, &util.InterfaceTypeErr{Name: reflect.TypeOf(roleDescOrType).String()}
	}
}

// CheckUserPermission 检查用户权限
func (rbac RBACAuth) CheckUserPermission(user model.User, permissionDescOrType interface{}) (bool, error) {
	permissions, err := rbac.GetUserPermissions(user.ID)
	if err != nil {
		return false, err
	}

	switch permissionDescOrType.(type) {
	case uint8:
		for _, p := range permissions {
			if p.Type == permissionDescOrType {
				return true, nil
			}
		}
		return false, nil
	case string:
		for _, p := range permissions {
			if p.Desc == permissionDescOrType {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, &util.InterfaceTypeErr{Name: reflect.TypeOf(permissionDescOrType).String()}
	}
}
