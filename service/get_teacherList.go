package service

import (
	"PingLeMe-Backend/serializer"
)

type GetTeacherListService struct {

}

func (service *GetTeacherListService) GetTeacherList() serializer.Response {

	return serializer.BuildTeacherListResponse()
}