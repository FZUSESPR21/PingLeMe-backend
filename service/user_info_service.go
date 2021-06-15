package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"errors"

	"gorm.io/gorm"
)

type UserInfoService struct {
	model.UserRepositoryInterface
	model.PairRepositoryInterface
}

// Information 根据userID获取学生、助教、老师信息
func (service *UserInfoService) Information(userID uint) serializer.Response {
	user, err := service.GetUser(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.ParamErr("user not found", err)
		} else {
			return serializer.ParamErr("", err)
		}
	}
	switch user.Role {
	case model.RoleStudent:
		pairID, err := service.GetPairByStudentID(user.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				pairID = 0
			} else {
				pairID = 0
			}
		}

		var pair model.User
		if pairID != 0 {
			pair, err = service.GetUser(pairID)
			if err != nil {
				pair = model.User{UID: "0", UserName: ""}
			}
		}

		teamID, err2 := service.GetUserTeamID(user)
		if err2 != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				teamID = 0
			} else {
				teamID = 0
				//return serializer.DBErr("", err2)
			}
		}
		classID, err3 := service.GetStudentClassID(user.ID)
		if err3 != nil {
			return serializer.ServerInnerErr("student has no class", err3)
		}
		return serializer.BuildStudentResponse(user, pair.UID, pair.UserName, teamID, classID)

	case model.RoleTeacher:
		return serializer.BuildTeacherResponse(user)

	default:
		return serializer.BuildUserResponse(user)
	}
}
