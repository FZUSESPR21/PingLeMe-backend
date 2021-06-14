//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// User 用户序列化器
type User struct {
	UID      string `json:"uid"`
	UserName string `json:"user_name"`
	PairUID  string `json:"pair_uid"`
	PairName string `json:"pair_name"`
	TeamID   uint   `json:"team_id"`
	ClassID  uint   `json:"class_id"`
	Role     uint8  `json:"role"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		UID:      user.UID,
		UserName: user.UserName,
	}
}

// BuildStudent 序列化学生
func BuildStudent(user model.User, pairUID, pairName string, teamID uint, classID uint) User {
	return User{
		UID:      user.UID,
		UserName: user.UserName,
		PairUID:  pairUID,
		PairName: pairName,
		ClassID:  classID,
		TeamID:   teamID,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

// BuildStudentResponse 序列化学生响应
func BuildStudentResponse(user model.User, pairUID, pairName string, teamID uint, classID uint) Response {
	return Response{
		Data: BuildStudent(user, pairUID, pairName, teamID, classID),
	}
}
