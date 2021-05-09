package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"gorm.io/gorm"
)

type TeammateSetService struct {
	model.TeamRepositoryInterface
	gorm.Model
	Students []model.User `form:"students" json:"students"`
}

func (service *TeammateSetService) AddTeammate() serializer.Response {
	if has, err := service.SetTeammate(service.ID, service.Students); err != nil || has != 1 {
		return serializer.DBErr("数据获取错误", err)
	}
	//TODO 相同名称判断
	//TODO 对学生身份判断和是否已有团队判断
	return serializer.Response{
		Code: 0,
		Msg:  "添加组员成功",
	}
}

func (service *TeammateSetService) DeleteTeammate() serializer.Response {
	if has, err := service.SetTeammate(service.ID, service.Students); err != nil || has != 1 {
		return serializer.DBErr("数据出现错误", err)
	}
	//TODO 相同名称判断
	return serializer.Response{
		Code: 0,
		Msg:  "删除组员成功",
	}
}
