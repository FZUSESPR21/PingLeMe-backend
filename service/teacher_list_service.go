package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type TeacherListService struct {
	model.UserRepositoryInterface
}

func (service *TeacherListService) GetTeacherList() serializer.Response {
	has, teacherList, err := service.GetAllTeacher()
	if err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	if has == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "目前没有老师！",
		}
	}

	return serializer.Response{
		Code: 0,
		Data: teacherList,
		Msg:  "",
	}
}

func (service *TeacherListService) GetAssistantList() serializer.Response {
	var assistantList []model.Assistant
	var err error
	var has int64
	if has, assistantList, err = service.GetAllAssistant(); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	if has == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "目前没有助教！",
		}
	}

	return serializer.Response{
		Code: 0,
		Data: assistantList,
		Msg:  "",
	}
}
