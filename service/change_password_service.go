//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"
	"strings"
)

// ChangePasswordService 修改密码的服务
type ChangePasswordService struct {
	model.UserRepositoryInterface
	UID      			 string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	OldPassword 		 string `form:"old_password" json:"old_password" binding:"required,min=8,max=40"`
	NewPassword 		 string `form:"new_password" json:"new_password" binding:"required,min=8,max=40"`
	NewPasswordConfirm	 string `form:"new_password_confirm	" json:"new_password_confirm" binding:"required,min=8,max=40"`
}

// ChangePassword 修改密码函数
func (service *ChangePasswordService) ChangePassword() serializer.Response {
	user, err := service.GetUserByUID(service.UID)
	if err != nil {
		return serializer.ParamErr("该用户不存在", err)
	}

	if user.CheckPassword(service.OldPassword) == false {
		err = errors.New("旧密码错误")
		return serializer.ParamErr("旧密码错误", err)
	}

	if strings.Compare(service.NewPassword, service.NewPasswordConfirm) != 0 {
		err = errors.New("两次密码不一致")
		return serializer.ParamErr("两次密码不一致", err)
	}

	err = user.SetPassword(service.NewPassword)
	if err != nil {
		return serializer.ParamErr("密码设置错误", err)
	}

	err = service.ChangeUserPassword(user, user.PasswordDigest)
	if err != nil {
		return serializer.DBErr("修改密码失败", err)
	}

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildUser(user),
		Msg:  "Success",
	}
}