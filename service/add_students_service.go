package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type AddStudentsService struct {
	model.UserRepositoryInterface
	model.ClassRepositoryInterface
	Students []StuInfo `form:"students" json:"students"`
	//TODO 接口需要修改，不是数组
}

type StuInfo struct {
	UID      string `form:"uid" json:"uid"`
	Nickname string `form:"name" json:"name"`
	ClassId  int    `form:"class_id" json:"class_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	//TODO ClassID 没地方存
}

func transformStruct(stuInfo StuInfo) (model.User, error) {
	var user model.User
	user.UID = stuInfo.UID
	user.UserName = stuInfo.Nickname
	err := user.SetPassword(stuInfo.Password)
	return user, err
}

//TODO UID相同判断,但接口中没有要求
func (service *AddStudentsService) AddStudents() serializer.Response {
	var user []model.User
	for _, a := range service.Students {
		u, err := transformStruct(a)
		if err!= nil {
			return serializer.ParamErr("", err)
		}
		user = append(user, u)
	}

	if err := service.SetUsers(user); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}

	for i, a := range user {
		class, err1 := service.GetClassByID(service.Students[i].ClassId)
		if err1 != nil {
			return serializer.ParamErr("该班级不存在", err1)
		}
		err1 = service.AddStudent(class, a)
		if err1 != nil {
			return serializer.DBErr("添加学生失败", err1)
		}
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
