package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type SetClassInfoService struct {
	model.ClassRepositoryInterface
	Name  string `form:"class_name" json:"class_name"`
}
//TODO 接口缺失 form:"class_name"
func (service *SetClassInfoService) SetClassInfo() serializer.Response {
	if has,err := service.GetClassByName(service.Name); err != nil {
		return serializer.DBErr("数据获取错误", err)
	} else if has == 1 {
		return serializer.DBErr("班级名称已存在！", err)
	}
	if err := service.UpdateClassName(service.Name); err != nil {
		return serializer.DBErr("数据获取错误", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

