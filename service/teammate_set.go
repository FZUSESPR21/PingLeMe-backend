package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"gorm.io/gorm"
)

type TeammateSetService struct {
	model.TeamRepositoryInterface
	model.UserRepositoryInterface
	gorm.Model
	UID        string `form:"uid" json:"uid"`
	TeamNumber int    `form:"teamNumber" json:"teamNumber"`
}

func makeStu(uid string) model.User {
	var student model.User
	student.UID = uid
	return student
}

func (service *TeammateSetService) AddTeammate() serializer.Response {
	//TODO 1.对学生身份判断 2.是否已有团队判断 3.团队是否存在
	var has int64
	var err error
	var user model.User
	user, err = service.GetUserByUID(service.UID)
	if err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	if has, err = service.AddTeammateByID(int(user.ID), service.TeamNumber); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	if has != 1 {
		return serializer.DBErr("has != 1 错误", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "添加组员成功",
	}

}

func (service *TeammateSetService) DeleteTeammate() serializer.Response {
	//TODO 1.对学生身份判断 2.是否已有团队判断 3.团队是否存在
	var has int64
	var err error
	var user model.User
	user, err = service.GetUserByUID(service.UID)
	if err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	if has, err = service.DeleteTeammateByID(int(user.ID)); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	if has != 1 {
		return serializer.DBErr("has != 1 错误", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "删除组员成功",
	}
}
