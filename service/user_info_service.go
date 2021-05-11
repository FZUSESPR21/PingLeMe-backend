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

// Information 学生、助教、老师信息
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
				return serializer.DBErr("", err)
			}
		}

		var pair model.User
		if pairID != 0 {
			pair, err = service.GetUser(pairID)
			if err != nil{
				return serializer.ParamErr("", err)
			}
		}

		teamID, err2 := service.GetUserTeamID(user)
		if err2 != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				teamID = 0
			} else {
				return serializer.DBErr("", err2)
			}
		}
		return serializer.BuildStudentResponse(user, pair.UID, pair.UserName, teamID)
	default:
		return serializer.BuildUserResponse(user)
	}
}
