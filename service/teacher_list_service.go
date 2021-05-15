package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type TeacherListService struct {
	model.UserRepositoryInterface
}

func (service *TeacherListService) GetTeacherList() serializer.Response {
	var teacherList []model.User
	var err error
	var has int64
	if has, teacherList, err = service.GetAllTeacher(); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	if has == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "目前没有老师！",
		}
	}

	return serializer.BuildTeacherListResponse(teacherList)
}
