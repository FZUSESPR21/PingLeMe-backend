//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	UserNick  string `json:"user_nick"`
	PairStatus  string `json:"pair_status"`
	TeamStatus  string `json:"team_status"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	if (user.Role != 0){
		return User{
			ID:         user.ID,
			UserName:	user.UID,
			UserNick:	user.Nickname,
			Password  :	user.PasswordDigest,
			CreatedAt: 	user.CreatedAt.Unix(),
		}
	} else {
		return User{
			ID:         user.ID,
			UserName:	user.UID,
			UserNick:	user.Nickname,
			PairStatus:	user.PairStatus,
			TeamStatus:	user.TeamStatus,
			Password  :	user.PasswordDigest,
			CreatedAt: 	user.CreatedAt.Unix(),
		}
	}

}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User)Response {
	return Response{
		Data: BuildUser(user),
	}

}
