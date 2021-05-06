//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// User 用户序列化器
type User struct {
	UID       uint   `json:"uid"`
	UserName  string `json:"user_name"`
	Role 	  uint8  `json:"role"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		UID:       user.ID,
		UserName:  user.UID,
		Role: 	   user.Role,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
