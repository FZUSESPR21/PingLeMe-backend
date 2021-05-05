package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type AddStudentService struct {
	model.UserRepositoryInterface
	UserList []StuInfo `form:"userList" json:"userList"`
	//TODO 接口需要修改，不是数组
}

type StuInfo struct {
	UID            string `form:"uid" json:"uid"`
	Nickname       string `form:"nickname" json:"nickname"`
	ClassID        string `form:"classid" json:"classid"`
	//TODO ClassID 没地方存
}

func transformStruct(stuInfo StuInfo) model.User {
	var user model.User
	user.UID = stuInfo.UID
	user.Nickname = stuInfo.Nickname
	return user
}

//TODO UID相同判断,但接口中没有要求
func (service *AddStudentService) AddStudent() serializer.Response {
	var length = len(service.UserList)
	var user []model.User
	for i := 0; i < length; i++ {
		a:=service.UserList[i]
		user = append(user, transformStruct(a))
	}
	if err := service.SetUsers(user); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}