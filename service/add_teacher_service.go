package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"fmt"
)

type AddTeacherService struct {
	model.UserRepositoryInterface
	Teachers []TeacherInfo `json:"teachers"` //结构体怎么命名
	//TODO 接口需要修改，不是数组
}

//添加时重复名称能否通过数据库检查
type TeacherInfo struct {
	UID      string `form:"uid" json:"uid"`
	Password string `form:"password" json:"password"`
}

func transformTeacher(teacherInfo TeacherInfo) model.User {
	var user model.User
	user.UID = teacherInfo.UID
	user.PasswordDigest = teacherInfo.Password
	return user
}

func (service *AddTeacherService) AddTeacher() serializer.Response {
	var errMes string
	var length = len(service.Teachers)
	//fmt.Println(length)
	for i := 0; i < length; i++ {
		a := service.Teachers[i]
		if _, err := service.AddTeacherByUser(transformTeacher(a)); err != nil {
			errMes += "账号" + a.UID + "添加时发生错误！"
			fmt.Println(err)
		}
	}
	if len(errMes) == 0 {
		errMes += "Success"
	}
	return serializer.Response{
		Code: 0,
		Msg:  errMes,
	}
}
