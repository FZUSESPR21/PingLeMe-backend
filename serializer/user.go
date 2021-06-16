//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
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
		ID:       user.ID,
		UID:      user.UID,
		UserName: user.UserName,
		Role:     user.Role,
	}
}

// BuildStudent 序列化学生
func BuildStudent(user model.User, pairUID, pairName string, teamID uint, classID uint) User {
	return User{
		ID:       user.ID,
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

// BuildStudentListResponse 序列化学生列表
func BuildStudentListResponse(user []model.User, classID int) Response {
	studentList := make([]User, 0)

	for _, u := range user {
		studentList = append(studentList, BuildStudent(u, "", "", 0, uint(classID)))
	}

	return Response{
		Code: 0,
		Data: studentList,
	}
}
