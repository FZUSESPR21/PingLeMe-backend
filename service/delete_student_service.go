package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"gorm.io/gorm"
)

type DeleteStudentService struct {
	gorm.Model
	model.UserRepositoryInterface
	UID string `form:"uid" json:"uid"`
}

func (service *DeleteStudentService) DeleteStudent() serializer.Response {
	if has, err := service.GetUserByUID(service.UID); err != nil {
		return serializer.DBErr("数据获取错误", err)
	} else if has.UID == "" {
		return serializer.DBErr("该账号不存在！", err)
	}
	if err := service.DeleteUser(service.ID); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
