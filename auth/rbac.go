//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package auth

import (
	"PingLeMe-Backend/model"
)

type RBACAuth struct {
	model.RBACRepositoryInterface
}

// CheckUserRole 检查用户角色
func (rbac RBACAuth) CheckUserRole(user model.User, roleDesc string) (bool, error) {
	roles, err := rbac.GetUserRoles(user.ID)
	if err != nil {
		return false, err
	}

	for _, r := range roles {
		if r.Desc == roleDesc {
			return true, nil
		}
	}
	return false, nil
}

// CheckUserPermission 检查用户权限
func (rbac RBACAuth) CheckUserPermission(user model.User, permissionDesc string) (bool, error) {
	permissions, err := rbac.GetUserPermissions(user.ID)
	if err != nil {
		return false, err
	}
	for _, p := range permissions {
		if p.Desc == permissionDesc {
			return true, nil
		}
	}
	return false, nil
}
